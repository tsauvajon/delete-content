package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	total := flag.Int("nb", 0, "number of files to delete (0 to delete every file)")
	workers := flag.Int("w", 10, "number of concurrent workers")
	flag.Parse()
	dirs := flag.Args()

	fmt.Printf("deleting %d folders with %d concurrent workers\n", len(dirs), (*workers))

	for _, dir := range dirs {
		paths, err := filepath.Glob(dir + "/*")
		if err != nil {
			panic(err)
		}

		// use a "local" total for the current dir, that can be reassigned
		total := (*total)

		if total > len(paths) || total == 0 {
			total = len(paths)
		}

		// fmt.Printf("deleting %d out of %d files in %s with %d workers\n", total, len(paths), dir, (*workers))
		fmt.Print(dir)
		fmt.Print(" ")

		files := make(chan string, total)
		done := make(chan bool, total)

		for i := 0; i < (*workers); i++ {
			go worker(files, done)
		}

		for i := 0; i < total; i++ {
			files <- paths[i]
		}
		close(files)

		loading := (total / 10)
		if loading == 0 {
			loading = 1
		}

		for i := 0; i < total; i++ {
			<-done
			if i%loading == 0 {
				fmt.Print("#")
			}
		}

		fmt.Println(" done")
	}
}

func worker(files <-chan string, done chan<- bool) {
	for f := range files {
		if err := os.Remove(f); err != nil {
			done <- false
			continue
		}
		done <- true
	}
}
