package fileutils

import (
	"os"
	"path/filepath"
)

func CopyFilesInDir(src, dst string) error {
	dir, err := os.Open(src)
	if err != nil {
		return err
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}
	for _, fileInfo := range fileInfos {
		srcPath := filepath.Join(src, fileInfo.Name())
		dstPath := filepath.Join(dst, fileInfo.Name())
		if fileInfo.IsDir() {
			continue
		}
		err := CopyFile(srcPath, dstPath)
		if err != nil {
			return err
		}
	}
	return nil
}

// Copy all files in src into dst (cp src/*)
// func CopyFilesInDir(src, dst string) error {
// 	_, err := exec.Command("cp", "-r", src, dst).Output()
// 	return err
// }
// find /path/to/source -type f -exec cp {} /path/to/destination \;
// find tests -type f -exec cp {} foo \;
