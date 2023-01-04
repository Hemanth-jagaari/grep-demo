package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(args) < 3 {
		log.Fatal("Enter Required Parameters")
	}
	var pattern string
	var flags []string
	var files []string

	n := len(args)
	var index = -1

	index = GetIndex(args, index)
	if index == -1 {
		log.Fatal("no pattern found")
	}
	pattern = args[index]
	flags = AddItems(1, index, args)
	files = AddItems(index+1, n, args)
	if pattern == "" || pattern == " " {
		log.Fatal("Select Other Pattern")
	}
	if len(files) == 0 {
		log.Fatal("Input Any file Name")
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
