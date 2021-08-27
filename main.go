package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const ZoneId = ":Zone.Identifier"

type fileSearcher func(path string) ([]string, error)

func recursiveFileSearch(path string) ([]string, error) {
	files := make([]string, 0)

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, path)
			return nil
		})
	return files, err
}

func fileRemover(dir string, fn fileSearcher) error {
	files, err := fn(dir)
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
		fileRemover(dir, recursiveFileSearch)
	}

	fmt.Println("Done! 🍰")
}
