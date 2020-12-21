package helper

import (
	"fmt"
	"io"
)

// StreamOutputBuffer returns the next buffer output from cmdOutput as a string.
// Will wait until buffer is written too.
func StreamOutputBuffer(cmdOutput io.ReadCloser, bufferSize int) string {
	buffer := make([]byte, bufferSize)
	n, err := cmdOutput.Read(buffer)

	if n > 0 {
		outputBuffer := string([]byte(buffer[0:n]))
		return outputBuffer
	}

	if err != nil {
		fmt.Println("Logger has stopped", err)
		errMessage := "Application has exited"
		return errMessage
	}

	return ""
}
