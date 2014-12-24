package main

import (
	"../../photorg"
	"../../photorg/pathtools"
	"fmt"
	"log"
)

func main() {
	// fname := "2014-12-24 06.37.44-7.jpg"
	// fpath := "/data/Dropbox/Photos/Photostream/2014/"

	// processJpegFile(fpath + fname)

	fmt.Println("Photo Reorganizer")

	options := photorg.Options{
		SourcePath: "/data/Dropbox/Camera Uploads/",
		DestRoot:   "/data/Dropbox/Photos/Photostream/",
	}

	// Register override
	photorg.RegisterDecoder("zip", decodeDateTakenFromLStat)

	info, err := photorg.OrganizeFiles(options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Files Moved:\n", info.NumFilesMoved)
	for _, f := range info.FilesMoved {
		fmt.Println(" -", f.DestPath)
	}
	fmt.Printf("%d Files Errored:\n", info.NumFilesError)
	for _, f := range info.FilesError {
		fmt.Println(" -", f)
		// fmt.Printf(" - >%s<\n", f)
	}
}

func decodeDateTakenFromLStat(moveInfo *photorg.MoveInfo) error {
	moveInfo.DateTaken = pathtools.ModTime(moveInfo.SourcePath)
	return nil
}
