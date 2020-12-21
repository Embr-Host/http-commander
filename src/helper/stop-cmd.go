package helper

import (
	"os/exec"
)

// StopCmd checks if cmd is running. If so it stops the process.
func StopCmd(cmd *exec.Cmd) {
	if cmd.ProcessState == nil {
		cmd.Process.Kill()
		cmd.Wait()
	}
}
