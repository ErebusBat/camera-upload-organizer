package photorg

import (
	"fmt"
	"regexp"
	"time"
)

func init() {
	handledExts := []string{
		".mov",
		".png",
	}
	registerSystemDecoders(handledExts, 10, "Filename", decodeDateTakenFromFileName)
}

// var reFilenameDateDecode *regexp.Regexp = regexp.MustCompile(`(?i)^(\d{4})-(\d{1,2})-(\d{1,2})[ 0-9\.tz]+\.[a-z]{1,5}$`)
var reFilenameDateDecode *regexp.Regexp = regexp.MustCompile(`(?i)^(\d{4})-(\d{1,2})-(\d{1,2}) (\d{2})\.(\d{2})\.(\d{2}).[a-z]{1,5}$`)

// Takes a filePath and attempts to parse dateTime from filename
// If success then returns true
func decodeDateTakenFromFileName(moveInfo *MoveInfo) error {

	matches := reFilenameDateDecode.FindStringSubmatch(moveInfo.fileName)
	if matches == nil {
		return fmt.Errorf("Could not decode date/time from >%s<", moveInfo.fileName)
	}

	dateString := matches[1] + "-" +
		matches[2] + "-" +
		matches[3] + " " +
		matches[4] + ":" +
		matches[5] + ":" +
		matches[6]

	localTZ := time.Now().Local().Location()
	dateTaken, err := time.ParseInLocation("2006-01-02 15:04:05", dateString, localTZ)
	if err != nil {
		fmt.Println(err)
		return err
	}
	moveInfo.DateTaken = &dateTaken

	return nil
}
