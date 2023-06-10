package fileutils

import "os/exec"

func CopyDir(src, dst string) error {
	_, err := exec.Command("cp", "-r", src, dst).Output()
	return err
}
