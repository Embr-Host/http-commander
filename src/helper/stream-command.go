package helper

import (
	"fmt"
	"io"
	"os/exec"
)

// StreamCommand returns the StdoutPipe of the cmd.
func StreamCommand(cmd *exec.Cmd) io.ReadCloser {
	cmdOutput, errOut := cmd.StdoutPipe()

	if errOut != nil {
		fmt.Println("Error creating StdoutPipe for Cmd", errOut)
	}

	return cmdOutput
}
