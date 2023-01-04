package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	var pattern string
	var flags []string
	var files []string

	i := 1
	n := len(args)
	var index = -1
	for i < n {
		if strings.Compare(args[i][:1], "-") != 0 {
			index = i
			break
		}
		i++
	}

	i = 1
	if index == -1 {
		log.Fatal("no pattern found")
	}
	pattern = args[index]
	for i < index {
		flags = append(flags, args[i])
		i++
	}
	i++
	for i < n {
		files = append(files, args[i])
		i++
	}
	patternList := GetPatterns(pattern)
	pattern = strings.Join(patternList, "|")
	NumberOfFiles = len(files)
	files = GetFilePaths(files)
	finalList := Search(pattern, flags, files)
	for _, line := range finalList {
		fmt.Printf("%s\n", line)
	}
}
