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
	total := flag.Int("nb", 0, "number of files to delete (0 to delete every file)")
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

		// use a "local" total for the current dir, that can be reassigned locally
		total := (*total)

		if total > len(paths) || total == 0 {
			total = len(paths)
		}

		fmt.Print(dir)
		fmt.Print(" ")

		files := make(chan string, total)
		status := make(chan int, len(paths))
		done := make(chan bool, 1)

		for i := 0; i < (*workers); i++ {
			go worker(files, done, status, limit)
		}

		loading := (total / 10)
		if loading == 0 {
			loading = 1
		}

		errors := 0
		skipped := 0
		count := 0

		finished := false

		var currentFileIndex int
		for currentFileIndex = 0; currentFileIndex < (*workers); currentFileIndex++ {
			files <- paths[currentFileIndex]
		}

		for !finished {
			switch <-status {
			case StatusError:
				errors++
			case StatusSkipped:
				skipped++
			case StatusOk:
				count++
				if count >= total {
					done <- true
					finished = true
				}
			}

			currentFileIndex++
			if currentFileIndex >= total {
				done <- true
				finished = true
			} else {
				files <- paths[currentFileIndex]
			}

			if count%loading == 0 {
				fmt.Print("#")
			}
		}

		close(files)
		close(status)

		fmt.Print(" done")
		if errors > 0 {
			fmt.Printf(" (%d errors)", errors)
		}
		if skipped > 0 {
			fmt.Printf(" (%d skipped)", skipped)
		}
		fmt.Println()
	}
}

func worker(files <-chan string, done <-chan bool, status chan<- int, limit time.Time) {
	for {
		var f string
		select {
		case f = <-files:
		case <-done:
			return
		}
		fi, err := os.Stat(f)
		if err != nil {
			status <- StatusError
			continue
		}

		mtime := fi.ModTime()

		if mtime.After(limit) {
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
