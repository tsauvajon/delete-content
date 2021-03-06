package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

const (
	// StatusOk : Deletion OK
	StatusOk = 0
	// StatusError : Deletion failed
	StatusError = 1
	// StatusSkipped : The file is too recent
	StatusSkipped = 2
)

func main() {
	daysLimit := flag.Int("d", 0, "only delete documents older than this number of days")
	workers := flag.Int("w", 10, "number of concurrent workers")
	flag.Parse()
	dirs := flag.Args()

	if len(dirs) == 0 {
		fmt.Println("at least 1 folder to delete")
		return
	}

	if *workers == 0 {
		fmt.Println("at least 1 worker")
		return
	}

	day := 24 * time.Hour
	days := time.Duration(-1*(*daysLimit)) * day
	limit := time.Now().Add(days)

	fmt.Printf("deleting %d folders with %d concurrent workers\n", len(dirs), (*workers))

	for _, dir := range dirs {
		paths, err := filepath.Glob(dir + "/*")
		if err != nil {
			panic(err)
		}

		fmt.Printf("deleting %d files in '%s' ", len(paths), dir)

		files := make(chan string, len(paths))
		status := make(chan int, len(paths))

		loading := (len(paths) / 10)
		if loading == 0 {
			loading = 1
		}

		errors := 0
		skipped := 0
		success := 0
		count := 0

		for i := 0; i < (*workers); i++ {
			go worker(files, status, limit)
		}

		var currentFileIndex int
		for currentFileIndex = 0; currentFileIndex < (*workers) && currentFileIndex < len(paths); currentFileIndex++ {
			files <- paths[currentFileIndex]
		}

		for i := 0; i < len(paths); i++ {
			switch <-status {
			case StatusError:
				errors++
			case StatusSkipped:
				skipped++
			default:
				success++
			}

			if currentFileIndex < len(paths) {
				files <- paths[currentFileIndex]
				currentFileIndex++
			}

			count++
			if count%loading == 0 {
				fmt.Print("#")
			}
		}

		close(files)

		fmt.Printf(" successfully deleted %d files", success)
		if errors > 0 {
			fmt.Printf(" (%d errors)", errors)
		}
		if skipped > 0 {
			fmt.Printf(" (%d skipped)", skipped)
		}
		fmt.Println()
	}
}

func worker(files <-chan string, status chan<- int, limit time.Time) {
	for f := range files {
		// can't read the filo info ?
		fi, err := os.Stat(f)
		if err != nil {
			status <- StatusError
			continue
		}

		// modified too recently to delete ?
		if mtime := fi.ModTime(); mtime.After(limit) {
			status <- StatusSkipped
			continue
		}

		if err := os.Remove(f); err != nil {
			status <- StatusError
			continue
		}

		status <- StatusOk
	}
}
