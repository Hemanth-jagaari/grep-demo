package main

import (
	"log"
	"os"
	"path"
	"strings"
)

func IsValidFile(fileName string) bool {

	ext := path.Ext(fileName)
	if compare(ext, ".txt") || compare(ext, ".log") || compare(ext, ".go") || compare(ext, ".py") || compare(ext, ".java") {
		return true
	}
	return false
}
func compare(ext, res string) bool {
	return strings.Compare(ext, res) == 0
}
func IsValidPath(fPath string) bool {

	return false
}
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
func GetFileName(fpath string) string {
	n := len(fpath)
	end := n
	start := n - 1
	for start >= 0 {
		if strings.Compare(fpath[start:start+1], "\\") == 0 || strings.Compare(fpath[start:start+1], "/") == 0 {
			break
		}
		start--
	}
	return fpath[start+1 : end]
}
func GetIndex(args []string, index int) int {
	i := 1
	n := len(args)
	for i < n {
		if strings.Compare(args[i][:1], "-") != 0 {
			index = i
			break
		}
		i++
	}
	return index
}
func AddItems(start, end int, args []string) []string {
	var list []string
	for start < end {
		list = append(list, args[start])
		start++
	}
	return list
}
