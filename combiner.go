package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	dir := "./"
	outputFile := "combined.md"

	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer out.Close()

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if path == outputFile {
			return nil
		}
		if filepath.Ext(path) == ".md" {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			_, err = out.Write(content)
			if err != nil {
				return err
			}
			_, err = out.WriteString("\n\n")
			if err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Failed to walk through directory: %v", err)
	}

	log.Println("All .md files have been combined into", outputFile)
}
