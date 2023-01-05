package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Search(t *testing.T) {

	var test_flag = []string{"-n"}
	var test_files = []string{"C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/tsample.txt"}
	var test_pattern = "balloon"

	assert.Equal(t, []string{"1:denotes a balloon whose horizontal", "4:y-coordinates of the balloon", "5:Arrows can be shot up directly balloon"}, Search(test_pattern, test_flag, test_files), "Wrong Lines returned")
	test_flag = append(test_flag, "-i")
	assert.Equal(t, []string{"1:denotes a balloon whose horizontal", "4:y-coordinates of the balloon", "5:Arrows can be shot up directly balloon"}, Search("Balloon", test_flag, test_files), "Wrong Matching lines")
	test_flag = append(test_flag, "-v")
	ans := []string{"2:diameter stretches between xstart", "3:and xend. You do not know the exact", "6:vertically in the positive y-direction", "7:from different points along the x-axis"}
	assert.Equal(t, ans, Search(test_pattern, test_flag, test_files), " wrong matching lines")
	test_flag = append(test_flag, "-l")
	assert.Equal(t, []string(nil), Search(test_pattern, test_flag, test_files), " wrong matching lines")

	test_files = append(test_files, "C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/texample.txt")
	assert.Equal(t, []string{"texample.txt"}, Search("Balloon", test_flag, test_files), "Not Matching")

	test_flag = append(test_flag, "-x")
	test_files = append(test_files, "C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/thistory.txt")
	assert.Equal(t, []string{"tsample.txt", "texample.txt", "thistory.txt"}, Search("Balloon", test_flag, test_files), "Wrong file names")

}
