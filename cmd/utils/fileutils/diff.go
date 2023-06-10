package fileutils

import "os/exec"

func Diff(src, dst string) (string, error) {
	bytes, err := exec.Command("diff", src, dst).Output()
	return string(bytes), err
}
