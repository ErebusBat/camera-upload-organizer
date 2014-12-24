package photorg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"./pathtools"
)

// Main API Entry, used
func OrganizeFiles(opts Options) (*ProcessedInfo, error) {
	if !pathtools.IsDir(opts.SourcePath) {
		return nil, fmt.Errorf("Invalid directory >%s<", opts.SourcePath)
	}
	opts.processed = ProcessedInfo{
		FilesMoved: make([]MoveInfo, 0, 50),
		FilesError: make([]string, 0, 10),
	}

	// use a closure so we can access opts instance
	walkClosure := func(path string, info os.FileInfo, err error) error {
		return walkVisit(&opts, path, info, err)
	}
	filepath.Walk(opts.SourcePath, walkClosure)

	// Do counts for them, ease of API use
	opts.processed.NumFilesMoved = len(opts.processed.FilesMoved)
	opts.processed.NumFilesError = len(opts.processed.FilesError)
	return &opts.processed, nil
}

func BuildMoveInfo(path string) MoveInfo {
	fileDir, fileName := filepath.Split(path)
	moveInfo := MoveInfo{
		SourcePath: path,
		fileDir:    fileDir,
		fileName:   fileName,
		fileExt:    filepath.Ext(path),
	}
	return moveInfo
}

// filepath.Walk callback, not in main closure because of length
func walkVisit(opts *Options, path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	// Build a MoveInfo and check to see if it is an ignored file
	moveInfo := BuildMoveInfo(path)
	if isIgnoredFile(&moveInfo) {
		// No error (still want to traverse)
		return nil
	}

	// Get Decoder and call it
	dFunc, decoderRegistered := GetDecoder(moveInfo.fileExt)
	if decoderRegistered {
		dFunc(&moveInfo)
	}

	// Check to make sure we have a DateTaken
	if moveInfo.DateTaken != nil {
		pathSuffix := GetDateTimePathSuffix(&moveInfo)
		moveInfo.DestPath = filepath.Join(opts.DestRoot, pathSuffix)
		opts.processed.FilesMoved = append(opts.processed.FilesMoved, moveInfo)
	} else {
		opts.processed.FilesError = append(opts.processed.FilesError, path)
		// Used to dump char codes (Icon file)
		// if strings.HasPrefix(moveInfo.fileName, "Icon") {
		//  for _, c := range moveInfo.fileName {
		//    fmt.Printf(">%s< %d\n", string(c), c)
		//  }
		// }
	}

	// Remember we are a callback for filepath.Walk, so return no error
	return nil
}

func GetDateTimePathSuffix(moveInfo *MoveInfo) (pathSuffix string) {
	dateTaken := moveInfo.DateTaken
	pathParts := make([]string, 0, 3)

	// Year
	pathParts = append(pathParts, dateTaken.Format("2006"))

	//Month
	pathParts = append(pathParts, dateTaken.Format("2006-01 January"))

	// File
	pathParts = append(pathParts, moveInfo.fileName)

	return strings.Join(pathParts, string(os.PathSeparator))
}
