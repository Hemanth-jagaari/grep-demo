package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

var NumberOfFiles int = 0
var name string = "jagaari hemanth"
var Flagcheck map[string]bool

func GetPwd() string {
	wrk, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wrk
}

func GetPatterns(pattern string) []string {

	var start = 0
	var n = len(pattern)
	i := 0
	var lst []string
	for i < n-1 {
		if strings.Compare(pattern[i:i+2], "\\|") == 0 {
			substr := pattern[start:i]
			start = i + 2
			i++
			lst = append(lst, substr)
		}
		i++
	}
	return lst
}
func GetFilePaths(filenames []string) []string {
	var lst []string
	n := len(filenames)

	i := 0
	pwd := GetPwd()
	for i < n {
		wrd := filenames[i]
		if strings.Compare(wrd[:1], "C") == 0 {
			lst = append(lst, wrd)
		} else {
			wrd = pwd + "\\" + wrd
			lst = append(lst, wrd)
		}
		i++
	}
	return lst
}
func Search(pattern string, flags, filenames []string) []string {
	log.Printf("flags:%v\n", flags)
	log.Printf("pattrens:%s\n", pattern)
	log.Printf("filepaths:%v\n", filenames)
	SetFlagChecks(flags)
	return FindLines(pattern, filenames)

}
func SetFlagChecks(flags []string) {
	for _, flag := range flags {
		Flagcheck[flag] = true
	}
}
func FindLines(pat string, files []string) []string {

	var totalList []string
	for _, file := range files {
		fileName := GetFileName(file)
		onefileLines, check := GetLines(pat, file)
		if check == true {
			totalList = append(totalList, fileName)
		} else {
			totalList = append(totalList, onefileLines...)
		}
	}
	return totalList
}
func GetLines(pat string, file string) ([]string, bool) {

	fileName := GetFileName(file)
	var lst []string
	f, err := os.Open(file)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	defer f.Close()

	var lineNumber int = 1
	for scanner.Scan() {
		line := scanner.Text()
		if IsMatch(line, pat) {

			var lineslices []string

			if Flagcheck["-l"] == true {
				return []string{}, true
			}

			if NumberOfFiles > 1 {
				lineslices = append(lineslices, fileName)
			}
			if Flagcheck["-n"] == true {
				lineslices = append(lineslices, string(lineNumber))
			}
			newline := strings.Join(lineslices, ":")
			lst = append(lst, newline)
		}
		lineNumber++
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lst, false
}

func GetFileName(filepath string) string {
	n := len(filepath)
	end := n
	start := n - 1
	for start >= 0 {
		if strings.Compare(filepath[start:start+1], "\\") == 0 {
			break
		}
		start--
	}
	return filepath[start+1 : end]
}

func IsMatch(line, pat string) bool {
	check := false

	_, err := regexp.MatchString(pat, line)
	if err == nil {
		check = true
	}
	if Flagcheck["-v"] == true {
		if check == false {
			check = true
		}
	}
	return check
}
