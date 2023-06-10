package cmdutils

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
)

// Wrapper for exec.Cmd that streams stdout, stderr, and error and outputs them in a struct
func Run(command string, arg ...string) CmdOutput {
	if !CmdExists(command) {
		return CmdOutput{Err: fmt.Errorf("command '%s' does not exist", command)}
	}
	cmd := exec.Command(command, arg...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return CmdOutput{Err: err}
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return CmdOutput{Err: err}
	}
	output := CmdOutput{}
	go stream(stdout, &output.Stdout)
	go stream(stderr, &output.Stderr)
	if err := cmd.Start(); err != nil {
		return CmdOutput{Err: err}
	}
	output.Err = cmd.Wait()
	return output
}

func stream(rc io.ReadCloser, output *string) {
	r := bufio.NewReader(rc)
	line, err := r.ReadString('\n')
	for err == nil {
		fmt.Print(line)
		*output += line
		line, err = r.ReadString('\n')
	}
}
