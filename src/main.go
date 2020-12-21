package main

import (
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

			for cmd.ProcessState == nil {
				outputBuffer := helper.StreamOutputBuffer(cmdOutput, bufferSize)
				writeResponse := sseStart + outputBuffer + sseEnd
				w.Write([]byte(writeResponse))
				flusher.Flush()
			}
		})
	})

	http.HandleFunc("/send-command", func(w http.ResponseWriter, r *http.Request) {
		middleware.Authenticate(settings, w, r, func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			httpCommand := r.PostForm.Get("command")
			fmt.Println("HTTP COMMAND", httpCommand)

			cmdInput.Write([]byte(httpCommand + "\n"))

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
