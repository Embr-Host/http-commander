package helper

import (
	"fmt"
	"io"
	"os/exec"
)

// InputCommand returns the StdinPipe of the cmd.
func InputCommand(cmd *exec.Cmd) (cmdInput io.WriteCloser) {
	cmdInput, cmdErr := cmd.StdinPipe()

	if cmdErr != nil {
		fmt.Println("Error creating StdinPipe for Cmd", cmdErr)
	}

	return cmdInput
}
