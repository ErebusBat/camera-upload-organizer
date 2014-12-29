package photorg

import (
	. "./pathtools"
	"log"
	"os"
)

func moveFile(moveInfo *MoveInfo) error {
	log.Printf("Moving: %s\n", moveInfo.DestPath)
	if !IsDir(moveInfo.destDir) {
		// if err := os.MkdirAll(moveInfo.destDir, os.ModeDir+755); err != nil {
		if err := os.MkdirAll(moveInfo.destDir, 0755); err != nil {
			return err
		}
	}
	if err := os.Rename(moveInfo.SourcePath, moveInfo.DestPath); err != nil {
		return err
	}
	return nil
}
