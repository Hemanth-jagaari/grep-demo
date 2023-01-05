package main

import (
	"reflect"
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

func TestSingleFileLines(t *testing.T) {
	testCases := []struct {
		pat       string
		file      string
		flagcheck map[string]bool
		expected  []string
	}{
		{
			pat:       "balloon",
			file:      "C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/texample.txt",
			flagcheck: map[string]bool{"-n": true, "-v": false, "-i": false, "-x": false, "-l": false},
			expected:  []string(nil),
		},
		{
			pat:       "balloons",
			file:      "C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/texample.txt",
			flagcheck: map[string]bool{"-n": true, "-v": false, "-i": false, "-x": false, "-l": false},
			expected: []string{"2:balloons taped onto a flat",
				"4:The balloons are represented"},
		},
		{
			pat:       "Balloons",
			file:      "C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/texample.txt",
			flagcheck: map[string]bool{"-n": true, "-v": false, "-i": true, "-x": false, "-l": false},
			expected:  []string{"2:balloons taped onto a flat", "4:The balloons are represented"},
		},
		{
			pat:       "Balloons",
			file:      "C:\\Users\\Hemanth\\Desktop\\go-projects\\grep-demo\\files/texample.txt",
			flagcheck: map[string]bool{"-n": true, "-v": false, "-i": true, "-x": false, "-l": true},
			expected:  []string{"texample.txt"},
		},
	}
	for i, tc := range testCases {
		Flagcheck = tc.flagcheck
		actual := SingleFileLines(tc.pat, tc.file)
		if !reflect.DeepEqual(actual, tc.expected) {
			t.Errorf("index= %d For pat=%q and file=%q, expected %q but got %q", i, tc.pat, tc.file, tc.expected, actual)
		}
	}
}

func TestHandleFullLine(t *testing.T) {
	cases := []struct {
		line, pat string
		result    bool
	}{
		{"hello hello", "hello", true},
		{"world world", "world", true},
		{"hello world", "foo", false},
		{"foo bar baz", "bar", false},
		{"foo bar baz", "baz", false},
		{"foo bar baz", "foo", false},
	}

	for _, c := range cases {
		got := HandleFullLine(c.line, c.pat)
		if got != c.result {
			t.Errorf("HandleFullLine(%q, %q) == %t, want %t", c.line, c.pat, got, c.result)
		}
	}
}
func TestIsMatch(t *testing.T) {
	cases := []struct {
		line, pat string
		result    bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", true},
		{"hello world", "foo", false},
		{"foo bar baz", "bar", true},
		{"foo bar baz", "baz", true},
		{"foo bar baz", "foo", true},
		{"foo bar baz", "qux", false},
	}

	// Test without -v flag.
	Flagcheck = map[string]bool{}
	for _, c := range cases {
		got := IsMatch(c.line, c.pat)
		if got != c.result {
			t.Errorf("IsMatch(%q, %q) == %t, want %t", c.line, c.pat, got, c.result)
		}
	}

	// Test with -v flag.
	Flagcheck = map[string]bool{"-v": true}
	for _, c := range cases {
		got := IsMatch(c.line, c.pat)
		if got == c.result {
			t.Errorf("IsMatch(%q, %q) == %t, want %t", c.line, c.pat, got, !c.result)
		}
	}
}
func TestFindMatch(t *testing.T) {
	cases := []struct {
		line, pat string
		result    bool
	}{
		{"hello world", "hello", true},
		{"hello world", "world", true},
		{"hello world", "foo", false},
		{"foo bar baz", "bar", true},
		{"foo bar baz", "baz", true},
		{"foo bar baz", "foo", true},
		{"foo bar baz", "qux", false},
	}

	// Test without -x flag.
	Flagcheck = map[string]bool{}
	for _, c := range cases {
		got := findMatch(c.line, c.pat)
		if got != c.result {
			t.Errorf("findMatch(%q, %q) == %t, want %t", c.line, c.pat, got, c.result)
		}
	}

	// Test with -x flag.
	Flagcheck = map[string]bool{"-x": true}
	for _, c := range cases {
		got := findMatch(c.line, c.pat)
		if got != HandleFullLine(c.line, c.pat) {
			t.Errorf("findMatch(%q, %q) == %t, want %t", c.line, c.pat, got, c.line == c.pat)
		}
	}
}
