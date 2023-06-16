package fileutils

import "os/exec"

// Recursively copy a directory from src to dst
//
// Literally uses cp -r
func CopyDir(src, dst string) error {
	_, err := exec.Command("cp", "-r", src, dst).Output()
	return err
}
