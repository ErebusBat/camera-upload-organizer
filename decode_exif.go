package photorg

import (
	"github.com/rwcarlsen/goexif/exif"
	"os"
)

func init() {
	RegisterDecoder(".jpg", decodeDateTakenExif)
	RegisterDecoder(".jpeg", decodeDateTakenExif)
}

func decodeDateTakenExif(moveInfo *MoveInfo) error {
	// fpath, fname := filepath.Split(filePath)
	f, err := os.Open(moveInfo.SourcePath)
	if err != nil {
		return err
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return err
	}

	dateTaken, _ := x.DateTime()
	moveInfo.DateTaken = &dateTaken
	// pathSuffix := GetDateTimePathSuffix(dateTaken, fname)
	// applyFileSuffix(filePath, dateTaken, pathSuffix)
	return nil
}
