package helper

import (
	"io"
	"os/exec"
	"strings"
)

// CommandRunner builds out a command from a string.
func CommandRunner(commandString string) (*exec.Cmd, io.ReadCloser, io.WriteCloser) {
	commandSplit := strings.Split(commandString, " ")
	cmd := exec.Command(commandSplit[0], commandSplit[1:]...)

	cmdOutput := StreamCommand(cmd)
	cmdInput := InputCommand(cmd)

	cmd.Start()
	return cmd, cmdOutput, cmdInput
}
