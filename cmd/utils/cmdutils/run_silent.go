package cmdutils

import (
	"bytes"
	"fmt"
	"os/exec"
)

// Wrapper for exec.Cmd that outputs stdout, stderr, and error in a struct
func RunSilent(command string, arg ...string) CmdOutput {
	if !CmdExists(command) {
		return CmdOutput{Err: fmt.Errorf("command '%s' does not exist", command)}
	}
	cmd := exec.Command(command, arg...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return CmdOutput{
		Err:    err,
		Stderr: stderr.String(),
		Stdout: stdout.String(),
	}
}
