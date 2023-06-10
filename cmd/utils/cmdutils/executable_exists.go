package cmdutils

import "os/exec"

func CmdExists(name string) bool {
	_, err := exec.LookPath(name)
	return err == nil
}
