package main

import (
	"../../photorg"
	// "../../photorg/pathtools"
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
	photorg.RegisterDecoder("zip", "LStat", decodeDateTakenFromLStat)
	photorg.RegisterDecoder("mov", "LStat", decodeDateTakenFromLStat)
	photorg.RegisterDecoder("png", "LStat", decodeDateTakenFromLStat)
	photorg.RegisterDecoder("jpg", "LStat", decodeDateTakenFromLStat)

	// dumpDecoderInfo("jpg")
	// return

	info, err := photorg.OrganizeFiles(options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Files Moved:\n", info.NumFilesMoved)
	// for _, f := range info.FilesMoved {
	// 	fmt.Println(" -", f.DestPath)
	// }
	fmt.Printf("%d Files Errored:\n", info.NumFilesError)
	// for _, f := range info.FilesError {
	// 	fmt.Println(" -", f)
	// 	// fmt.Printf(" - >%s<\n", f)
	// }
}

func dumpDecoderInfo(ext string) {
	decoders, haveAny := photorg.GetDecoders(ext)
	if !haveAny {
		fmt.Printf("No decoders registered >%s<\n", ext)
		return
	}
	fmt.Println("Found", len(decoders), "decoders for", ext)
	for _, dec := range decoders {
		fmt.Printf(" - [%d] %s\n", dec.Priority, dec.Name)
	}
}

func decodeDateTakenFromLStat(moveInfo *photorg.MoveInfo) error {
	// moveInfo.DateTaken = pathtools.ModTime(moveInfo.SourcePath)
	// log.Printf("%s -> %s\n", moveInfo.SourcePath, moveInfo.DateTaken)
	return nil
}
