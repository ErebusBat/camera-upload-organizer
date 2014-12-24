package photorg

import (
	"github.com/rwcarlsen/goexif/exif"
	"os"
)

func init() {
	handledExts := []string{
		".jpg",
		".jpeg",
	}
	registerSystemDecoders(handledExts, 0, "Exif", decodeDateTakenExif)
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
