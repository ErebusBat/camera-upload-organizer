package photorg

import (
	"time"
)

type Options struct {
	SourcePath string
	DestRoot   string
	processed  ProcessedInfo
}

type ProcessedInfo struct {
	FilesMoved    []MoveInfo
	NumFilesMoved int
	FilesError    []string
	NumFilesError int
}

type MoveInfo struct {
	SourcePath string
	DateTaken  *time.Time
	destDir    string
	DestPath   string
	fileDir    string
	fileName   string
	fileExt    string
}
