package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetFileName(t *testing.T) {
	filepath := "C:/Users/Username/desktop/folder/example.txt"
	ans := "example.txt"
	res := GetFileName(filepath)
	assert.Equal(t, ans, res, "Wrong fileNmae")

	filepath = "/example.txt"
	res = GetFileName(filepath)
	assert.Equal(t, ans, res, "Wrong fileName")

}
func Test_GetPwd(t *testing.T) {
	pwd := GetPwd()
	real, _ := os.Getwd()
	assert.Equal(t, real, pwd, "wrong PWD")
}

func Test_GetEmptyFlags(t *testing.T) {
	var flags = map[string]bool{"-n": false, "-l": false, "-i": false, "-v": false, "-x": false}

	assert.Equal(t, flags, GetEmptyFlags(), "Wrong init of Flags")
}
func Test_GetPatterns(t *testing.T) {

	var pat = "pat1"
	ans := GetPatterns(pat)
	assert.Equal(t, []string{pat}, ans, "Patterns are not Equal")

	pat = "pat1\\|pat2\\|pat3\\|pat4"
	ans = GetPatterns(pat)
	assert.Equal(t, []string{"pat1", "pat2", "pat3", "pat4"}, ans, "Pattrens are not equal")
}
func Test_GetFilePaths(t *testing.T) {

	ans := GetFilePaths([]string{"C:\\User\\name\\a.txt", "sample.txt", "files\\ex.txt"})

	pwd := GetPwd()
	assert.Equal(t, []string{"C:\\User\\name\\a.txt", pwd + "\\" + "sample.txt", pwd + "\\" + "files\\ex.txt"}, ans, "FIlepaths are not equal")
}
func Test_GetIndex(t *testing.T) {

	ans := GetIndex([]string{"arg0", "pat1", "sample.txt", "sample2.log"}, -1)
	assert.Equal(t, 1, ans, "Wrong Index Value")
	assert.Equal(t, 3, GetIndex([]string{"arg0", "-n", "-v", "pat1", "sample.txt", "sample2.log"}, -1), "Wrong Index Value")
}
func Test_AddItems(t *testing.T) {
	ans := AddItems(1, 3, []string{"0", "1", "2", "3", "4", "5"})
	assert.Equal(t, []string{"1", "2"}, ans, "Adding items is not equal")
	assert.Equal(t, []string{}, AddItems(4, 3, []string{"0", "1", "2", "3", "4", "5"}), "Adding items is not equal")
}
func Test_IsValidFile(t *testing.T) {
	val := IsValidFile("sample.sql")
	assert.Equal(t, false, val, "File Format is Not same")
}
