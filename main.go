package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	dirs := os.Args[4:]

	loading, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid loading: ", err)
		return
	}

	total, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Invalid total: ", err)
		return
	}

	workers, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("Invalid workers: ", err)
		return
	}

	for _, dir := range dirs {
		paths, err := filepath.Glob(dir + "/*")
		if err != nil {
			panic(err)
		}

		if total > len(paths) {
			total = len(paths)
		}

		fmt.Printf("deleting %v out of %v files with %v workers\n", total, len(paths), workers)

		files := make(chan string, total)
		results := make(chan bool, total)

		for i := 0; i < workers; i++ {
			go worker(files, results, loading)
		}

		for i := 0; i < total; i++ {
			files <- paths[i]
		}
		close(files)

		for i := 0; i < total; i++ {
			<-results
		}
	}

	fmt.Println(" done")
}

func worker(files <-chan string, results chan<- bool, loading int) {
	i := 0
	for f := range files {
		if err := os.Remove(f); err != nil {
			results <- false
			continue
		}
		if i%loading == 0 {
			fmt.Print("#")
		}
		i++
		results <- true
	}
}
