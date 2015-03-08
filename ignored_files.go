package photorg

// Ignore byte, just using byte for membership
var ignoreFiles map[string]byte

func isIgnoredFile(moveInfo *MoveInfo) bool {
	// return (moveInfo.fileExt != ".mov")
	if ignoreFiles == nil {
		setupIgnoredFiles()
	}
	_, isIgnored := ignoreFiles[moveInfo.fileName]
	return isIgnored
}
func setupIgnoredFiles() {
	ignoreFiles = make(map[string]byte)
	defaultIgnores := []string{
		".DS_Store",
		"Icon" + string(13), // Icon file has a CR in it's name
		"config.toml",
	}
	for _, f := range defaultIgnores {
		ignoreFiles[f] = '0'
	}
}
