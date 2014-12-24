package pathtools

import (
	"os"
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
