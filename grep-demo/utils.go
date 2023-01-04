package main

import (
	"log"
	"os"
	"strings"
)

func GetPwd() string {
	wrk, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return wrk
}
func SetFlagChecks(flags []string) {
	for _, flag := range flags {
		Flagcheck[flag] = true
	}
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
	substr := pattern[start:]
	lst = append(lst, substr)
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
func GetFileName(filepath string) string {
	n := len(filepath)
	end := n
	start := n - 1
	for start >= 0 {
		if strings.Compare(filepath[start:start+1], "\\") == 0 || strings.Compare(filepath[start:start+1], "/") == 0 {
			break
		}
		start--
	}
	return filepath[start+1 : end]
}
