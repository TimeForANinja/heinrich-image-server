package main

import (
	"log"
	"os"
	"path/filepath"
)

const lenPathSep = len(string(os.PathSeparator))

// ImageEntry represents an image entry with a name and folder.
type ImageEntry struct {
	Name   string `json:"name"`
	Folder string `json:"folder,omitempty"`
}

var imageEntries = []ImageEntry{}

// walkDirectory is a function to walk a directory up to a specified depth level
func walkDirectory(basedir string, depth int) ([]ImageEntry, error) {
	var entries []ImageEntry

	err := filepath.Walk(basedir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			// Skip directories
			return nil
		}

		folder, file := filepath.Split(path)

		ie := ImageEntry{Name: file}

		subpath := folder[len(basedir)+lenPathSep:]
		subpathLength := len(segmnentDir(subpath))
		if subpathLength > depth {
			return nil
		} else if subpathLength != 0 {
			ie.Folder = subpath[:len(subpath)-lenPathSep]
		}

		// Append the image entry
		entries = append(entries, ie)

		return nil
	})

	return entries, err
}

func segmnentDir(dir string) []string {
	dir = filepath.Clean(dir)
	if dir == "." {
		return []string{}
	}

	dir, last := filepath.Split(dir)
	return append(segmnentDir(filepath.Clean(dir)), last)
}

// updateImageEntries is a function to update the imageEntries array.
func updateImageEntries() {
	// Walk through the IMAGE_DIR with a specified (max) depth of 2
	entries, err := walkDirectory(IMAGE_DIR, 1)
	if err != nil {
		log.Println("Error updating image entries:", err)
		return
	}

	// assign the new list of entries
	imageEntries = entries
}
