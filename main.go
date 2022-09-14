package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const ZoneId = ":Zone.Identifier"

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

		files, err := searchFilesRecursively(dir)
		if err != nil {
			log.Fatal(err)
		}

		removeFiles(files)
	}

	fmt.Println("Done! 🍰")
}

func searchFilesRecursively(path string) ([]string, error) {
	files := make([]string, 0)

	err := filepath.Walk(path,
		func(p string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			files = append(files, p)
			return nil
		})
	return files, err
}

func removeFiles(files []string) error {
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
