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
	// filepath.Walk(opts.SourcePath, walkVisit)
	walkClosure := func(path string, info os.FileInfo, err error) error {
		return walkVisit(&opts, path, info, err)
	}
	filepath.Walk(opts.SourcePath, walkClosure)

	// Do counts for them, ease of API use
	opts.processed.NumFilesMoved = len(opts.processed.FilesMoved)
	opts.processed.NumFilesError = len(opts.processed.FilesError)
	return &opts.processed, nil
}

func walkVisit(opts *Options, path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return nil
	}

	fileDir, fileName := filepath.Split(path)
	moveInfo := MoveInfo{
		SourcePath: path,
		fileDir:    fileDir,
		fileName:   fileName,
	}
	if isIgnoredFile(&moveInfo) {
		return nil
	}

	switch ext := filepath.Ext(path); {
	case ext == ".jpg",
		ext == ".jpeg":
		decodeDateTakenExif(&moveInfo)
	default:
		decodeDateTakenFromFileName(&moveInfo)
	}

	if moveInfo.DateTaken != nil {
		pathSuffix := GetDateTimePathSuffix(&moveInfo)
		moveInfo.DestPath = filepath.Join(opts.DestRoot, pathSuffix)
		opts.processed.FilesMoved = append(opts.processed.FilesMoved, moveInfo)
	} else {
		// Used to dump char codes (Icon file)
		// if strings.HasPrefix(moveInfo.fileName, "Icon") {
		// 	for _, c := range moveInfo.fileName {
		// 		fmt.Printf(">%s< %d\n", string(c), c)
		// 	}
		// }
		opts.processed.FilesError = append(opts.processed.FilesError, path)
	}
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
