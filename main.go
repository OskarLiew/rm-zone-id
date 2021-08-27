package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const ZoneId = ":Zone.Identifier"

func recursiveRemove(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file, ZoneId) {
			err := os.Remove(file)
			if err != nil {
				return err
			}
			fmt.Printf("Deleted %s\n", file)
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: No path given!")
		os.Exit(1)
	}
	dirs := os.Args[1:]

	for _, dir := range dirs {
		dir, err := filepath.Abs(dir)
		if err != nil {
			log.Fatal(err)
		}
		recursiveRemove(dir)
	}

	fmt.Println("Done! 🍰")
}
