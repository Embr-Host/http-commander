package main

import (
	"encoding/json"
	"fmt"
	"helper"
	"middleware"
	"net/http"
)

func main() {

	settings := helper.GetSettingsFromFile("settings.json")

	fmt.Println("Starting command...", settings.StartCommand)

	cmd, cmdOutput, cmdInput := helper.CommandRunner(settings.StartCommand)
	bufferSize := 80

	cmd.Start()

	sseStart := "data: "
	sseEnd := "\n\n"

	http.HandleFunc("/logging", func(w http.ResponseWriter, r *http.Request) {
		middleware.Authenticate(settings, w, r, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Transfer-Encoding", "chunked")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")

			flusher, ok := w.(http.Flusher)
			if !ok {
				http.Error(w, "Connection does not support streaming", http.StatusBadRequest)
				return
			}

			for cmd.ProcessState == nil && r.Context().Err() == nil {
				outputBuffer := helper.StreamOutputBuffer(cmdOutput, bufferSize)
				writeResponse := sseStart + outputBuffer + sseEnd
				if r.Context().Err() == nil {
					w.Write([]byte(writeResponse))
					flusher.Flush()
				}
			}

			hijacker := w.(http.Hijacker)
			conn, readBuffer, err := hijacker.Hijack()
			_ = readBuffer

			if err != nil {
				fmt.Println(err, "Unknown error has occurred... Attempting to close ideal connection")
				http.DefaultClient.CloseIdleConnections()
			}

			conn.Close()
		})
	})

	type CommandBody struct {
		Command string `json:"command"`
	}
	http.HandleFunc("/send-command", func(w http.ResponseWriter, r *http.Request) {
		middleware.Authenticate(settings, w, r, func(w http.ResponseWriter, r *http.Request) {
			var data CommandBody
			json.NewDecoder(r.Body).Decode(&data)
			fmt.Println("HTTP COMMAND", data.Command)

			cmdInput.Write([]byte(data.Command + "\n"))

			w.Write([]byte("OK!"))
		})
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		if cmd.ProcessState != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Internal Server Error"))
		} else {
			w.Write([]byte("OK!"))
		}

	})

	http.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
		middleware.Authenticate(settings, w, r, func(w http.ResponseWriter, r *http.Request) {
			helper.StopCmd(cmd)
			w.Write([]byte("OK!"))
		})
	})

	http.HandleFunc("/restart", func(w http.ResponseWriter, r *http.Request) {
		middleware.Authenticate(settings, w, r, func(w http.ResponseWriter, r *http.Request) {
			helper.StopCmd(cmd)
			cmd, cmdOutput, cmdInput = helper.CommandRunner(settings.StartCommand)

			w.Write([]byte("OK!"))
		})
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK!"))
	})

	fmt.Println("Starting server on port", settings.CommanderPort)
	http.ListenAndServe(":"+settings.CommanderPort, nil)
}
