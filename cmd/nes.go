package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nathanperkins/go-nes/nes"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "%s [flags] rom_filename", os.Args[0])
		flag.PrintDefaults()
		os.Exit(2)
	}

	if len(os.Args) < 2 {
		flag.Usage()
	}

	filename := os.Args[1]

	r, err := nes.NewNES(filename)
	if err != nil {
		log.Fatalf("Could not open ROM: %v", err)
	}

	if err := r.Run(); err != nil {
		log.Fatalf("NES exited early: %v", err)
	}
}
