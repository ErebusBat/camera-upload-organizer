package pathtools

import (
	"os"
	"time"
)

func IsDir(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return fi.IsDir()
}

func IsFile(path string) bool {
	fi, err := os.Lstat(path)
	if err != nil {
		return false
	}
	// We can't just do !IsDir() because if path points to a non existant file
	// then it would erroniously return true
	return !fi.IsDir()
}

func ModTime(path string) *time.Time {
	fi, err := os.Lstat(path)
	if err != nil {
		return nil
	}
	// We can't just do !IsDir() because if path points to a non existant file
	// then it would erroniously return true
	modTime := fi.ModTime()
	return &modTime
}
