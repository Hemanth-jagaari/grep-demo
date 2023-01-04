package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	/*

		+++++++++++++++++++++++++++++steps++++++++++++++++++++++++++++++++++++++++
		process the command line arguments
		store all file names

			file extensions
			file paths are relative or fully qualified
			find fullfile path for all givenn files

		get the pattern
		it may be contains multiple patterns
		it may need further processing
		flag
		keep track of all flags and there implication on the generated output
		funcs to openfile readfile and findmatch in files

	*/
	args := os.Args
	//fmt.Println("In main function")
	//fmt.Printf("Type of cmd args is %T\n", args)
	//fmt.Printf("cmd line arguments are %v\n", args)

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
	//log.Printf("Raw: flags %v\n", flags)
	//log.Printf("Raw: patterrn %s\n", pattern)
	//log.Printf("Raw: files %v\n", files)

	patternList := GetPatterns(pattern)
	pattern = strings.Join(patternList, "|")
	NumberOfFiles = len(files)
	files = GetFilePaths(files)
	finalList := Search(pattern, flags, files)
	for _, line := range finalList {
		fmt.Printf("%s\n", line)
	}
}

/*func readLines(filepath string) []string {

	f, err := os.Open(filepath)

	if err != nil {
		panic(err)
	}
	var lst []string
	scanner := bufio.NewScanner(f)
	defer f.Close()
	var i int = 1
	for scanner.Scan() {
		line := fmt.Sprintf("line-%d:%s\n", i, scanner.Text())
		fmt.Println(line)
		lst = append(lst, line)
		i++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lst
}
*/
