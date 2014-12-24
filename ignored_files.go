package photorg

import (
// "fmt"
)

// Ignore byte, just using byte for membership
var ignoreFiles map[string]byte

func isIgnoredFile(moveInfo *MoveInfo) bool {
	if ignoreFiles == nil {
		setupIgnoredFiles()
	}
	_, isIgnored := ignoreFiles[moveInfo.fileName]
	return isIgnored
}
func setupIgnoredFiles() {
	ignoreFiles = make(map[string]byte)
	defaultIgnores := []string{".DS_Store", "Icon" + string(13)}
	for _, f := range defaultIgnores {
		ignoreFiles[f] = '0'
	}
}
