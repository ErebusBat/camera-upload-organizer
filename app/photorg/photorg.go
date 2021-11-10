package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
	photorg "github.com/ErebusBat/camera-upload-organizer"
	"github.com/ErebusBat/camera-upload-organizer/pathtools"
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

	// Register Fallbacks
	for _, ext := range photorg.RegisteredExtensions() {
		photorg.RegisterDecoder(ext, "LStat", decodeDateTakenFromLStat)
	}

	// Debugging
	// photorg.DumpDecodersInfo()
	// return

	info, err := photorg.OrganizeFiles(*options)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d Files Moved:\n", info.NumFilesMoved)
	fmt.Printf("%d Files Errored:\n", info.NumFilesError)
	for _, f := range info.FilesError {
		fmt.Println(" -", f)
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

// Fallback to use the file system modification time if we don't have anything else
func decodeDateTakenFromLStat(moveInfo *photorg.MoveInfo) error {
	moveInfo.DateTaken = pathtools.ModTime(moveInfo.SourcePath)
	return nil
}
