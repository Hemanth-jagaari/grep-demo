package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var NumberOfFiles int = 0
var Flagcheck = map[string]bool{"-n": false, "-l": false, "-i": false, "-v": false, "-x": false}

func Search(pattern string, flags, filenames []string) []string {
	SetFlagChecks(flags)
	return FindLines(pattern, filenames)
}
func FindLines(pat string, files []string) []string {

	var totalList []string
	for _, file := range files {
		fileName := GetFileName(file)
		onefileLines, check := GetLines(pat, file)
		if check {
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
	MatchCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		if IsMatch(line, pat) {
			var lineslices []string

			if Flagcheck["-l"] && !Flagcheck["-v"] {
				return []string{}, true
			} else if Flagcheck["-l"] && Flagcheck["-v"] {
				MatchCount++
				lineNumber++
				continue
			}
			if NumberOfFiles > 1 {
				lineslices = append(lineslices, fileName)
			}
			if Flagcheck["-n"] {
				lineslices = append(lineslices, strconv.Itoa(lineNumber))
			}
			lineslices = append(lineslices, line)
			newline := strings.Join(lineslices, ":")
			lst = append(lst, newline)
		}
		lineNumber++
	}
	if Flagcheck["-v"] && Flagcheck["-l"] {
		if MatchCount == lineNumber-1 {
			lst = append(lst, fileName)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return lst, false
}
func IsMatch(line, pat string) bool {
	check := false

	if findMatch(line, pat) {
		check = true
	}
	if Flagcheck["-v"] {
		if check {
			check = false
		} else {
			check = true
		}
	}
	return check
}
func findMatch(line, pat string) bool {
	wordList := strings.Split(line, " ")
	if Flagcheck["-x"] {
		return HandleFullLine(line, pat)
	}
	for _, wrd := range wordList {

		if Flagcheck["-i"] {
			if strings.EqualFold(wrd, pat) {
				return true
			}
		} else {
			if strings.Compare(wrd, pat) == 0 {
				return true
			}
		}
	}
	return false
}
func HandleFullLine(line, pat string) bool {

	words := strings.Split(line, " ")

	n := len(words)

	count := 0
	for _, wrd := range words {
		if Flagcheck["-i"] {
			if strings.EqualFold(wrd, pat) {
				count++
			}
		} else {
			if strings.Compare(wrd, pat) == 0 {
				count++
			}
		}
	}
	return n == count
}
