package main

import "os"

func main() {
	argsWithoutProg := os.Args[1:]
	for _, dir := range argsWithoutProg {
		os.RemoveAll(dir)
		os.MkdirAll(dir, 0777)
	}
}
