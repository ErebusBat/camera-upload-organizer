package photorg

import (
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func init() {
	handledExts := []string{
		".jpg",
		".jpeg",
	}
	registerSystemDecoders(handledExts, 1, "Exif", decodeDateTakenExif)
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
	return nil
}
