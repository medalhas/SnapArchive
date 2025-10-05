package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// OverwriteBehavior defines how to handle file conflicts
type OverwriteBehavior int

const (
	Skip OverwriteBehavior = iota
	Overwrite
	Rename
)

func main() {
	var overwriteMode string
	flag.StringVar(&overwriteMode, "overwrite", "skip", "Behavior when file exists: 'skip', 'overwrite', or 'rename'")
	flag.Parse()

	// Validate overwrite mode
	var behavior OverwriteBehavior
	switch overwriteMode {
	case "skip":
		behavior = Skip
	case "overwrite":
		behavior = Overwrite
	case "rename":
		behavior = Rename
	default:
		fmt.Fprintf(os.Stderr, "Error: Invalid overwrite mode '%s'. Valid options: skip, overwrite, rename\n", overwriteMode)
		os.Exit(1)
	}

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <source_folder> <dest_folder>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	sourceDir := args[0]
	destDir := args[1]

	if _, err := os.Stat(sourceDir); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "Error: Source directory '%s' does not exist\n", sourceDir)
		os.Exit(1)
	}

	if err := os.MkdirAll(destDir, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating destination directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Organizing files from '%s' to '%s' (overwrite mode: %s)...\n", sourceDir, destDir, overwriteMode)

	if err := organizeFiles(sourceDir, destDir, behavior); err != nil {
		fmt.Fprintf(os.Stderr, "Error organizing files: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("File organization completed successfully!")
}

// organizeFiles recursively walks through the source directory and copies files
// to the destination directory organized by year and month
func organizeFiles(sourceDir, destDir string, behavior OverwriteBehavior) error {
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		modifiedTime := info.ModTime()
		year := modifiedTime.Format("2006")
		month := modifiedTime.Format("2006-01-02")

		destPath := filepath.Join(destDir, year, month)
		if err := os.MkdirAll(destPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %v", destPath, err)
		}

		fileName := info.Name()
		destFile := filepath.Join(destPath, fileName)

		// Handle file conflicts based on behavior
		if _, err := os.Stat(destFile); err == nil {
			switch behavior {
			case Skip:
				fmt.Printf("Skipped (file exists): %s -> %s\n", path, destFile)
				return nil
			case Overwrite:
				fmt.Printf("Overwriting: %s -> %s\n", path, destFile)
			case Rename:
				destFile = getUniqueFileName(destFile)
				fmt.Printf("Renamed and copied: %s -> %s\n", path, destFile)
			}
		} else {
			fmt.Printf("Copied: %s -> %s\n", path, destFile)
		}

		if err := copyFile(path, destFile); err != nil {
			return fmt.Errorf("failed to copy %s to %s: %v", path, destFile, err)
		}

		return nil
	})
}

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	sourceInfo, err := os.Stat(src)
	if err != nil {
		return err
	}

	return os.Chmod(dst, sourceInfo.Mode())
}

// getUniqueFileName ensures the destination file name is unique
// by appending a number if a file with the same name already exists
func getUniqueFileName(filePath string) string {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return filePath
	}

	ext := filepath.Ext(filePath)
	baseName := filePath[:len(filePath)-len(ext)]

	counter := 1
	for {
		newPath := fmt.Sprintf("%s_%d%s", baseName, counter, ext)
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			return newPath
		}
		counter++
	}
}