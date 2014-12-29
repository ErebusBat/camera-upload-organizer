package main

import (
	"../../photorg"
	"../../photorg/pathtools"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
)

func parseConfig(path string) (config *photorg.Options, err error) {
	config = new(photorg.Options)
	log.Printf("Loading config %s\n", path)
	if _, err = toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	log.Printf("  SourcePath = %s\n", config.SourcePath)
	log.Printf("    DestRoot = %s\n", config.DestRoot)
	return config, nil
}

func main() {
	flag.Parse()
	tomlFile := "./config.toml"
	if len(flag.Args()) > 1 {
		log.Fatal("USAGE: photorg [config_file.toml]")
	} else if len(flag.Args()) == 1 {
		tomlFile = flag.Args()[0]
	}

	fmt.Println("Photo Reorganizer")

	options, err := parseConfig(tomlFile)
	if err != nil {
		log.Fatal(err)
	}
	// options := photorg.Options{
	//   SourcePath: "/data/Dropbox/Camera Uploads/",
	//   DestRoot:   "/data/Dropbox/Photos/Photostream/",
	// }
	// options := photorg.Options{
	//  SourcePath: "/data/void/dropb/cam",
	//  DestRoot:   "/data/void/dropb/photostream",
	// }

	// Register override
	photorg.RegisterDecoder("mov", "LStat", decodeDateTakenFromLStat)
	photorg.RegisterDecoder("png", "LStat", decodeDateTakenFromLStat)
	photorg.RegisterDecoder("jpg", "LStat", decodeDateTakenFromLStat)

	// dumpDecoderInfo("jpg")
	// return

	info, err := photorg.OrganizeFiles(*options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Files Moved:\n", info.NumFilesMoved)
	// for _, f := range info.FilesMoved {
	// 	fmt.Println(" -", f.DestPath)
	// }
	fmt.Printf("%d Files Errored:\n", info.NumFilesError)
	for _, f := range info.FilesError {
		fmt.Println(" -", f)
		// fmt.Printf(" - >%s<\n", f)
	}
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
	moveInfo.DateTaken = pathtools.ModTime(moveInfo.SourcePath)
	// log.Printf("%s -> %s\n", moveInfo.SourcePath, moveInfo.DateTaken)
	return nil
}
