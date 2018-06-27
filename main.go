package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	dirs := os.Args[3:]
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
	for _, dir := range dirs {
		files, err := filepath.Glob(dir + "/*")
		if err != nil {
			panic(err)
		}
		for i, f := range files {
			if i%loading == 0 {
				fmt.Print("#")
			}
			if (i+1)%total == 0 {
				break
			}
			if err := os.Remove(f); err != nil {
				panic(err)
			}
		}
	}
}
