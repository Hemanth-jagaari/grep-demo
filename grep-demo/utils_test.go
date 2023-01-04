package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func Init() {
	log.Println("hello here in testing")
}

func TestGetFileName(t *testing.T) {
	filepath := "C:/Users/Username/desktop/folder/example.txt"
	ans := "example.txt"
	res := GetFileName(filepath)

	if strings.Compare(ans, res) != 0 {
		t.Errorf("ans is %s but got %s\n", ans, res)
	}
	filepath = "example.txt"
	res = GetFileName(filepath)
	if strings.Compare(ans, res) != 0 {
		t.Errorf("ans is %s but got %s\n", ans, res)
	}

}
func TestGetPwd(t *testing.T) {
	pwd := GetPwd()
	real, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	if strings.Compare(pwd, real) != 0 {
		t.Errorf("Required %s we got %s\n", real, pwd)
	}

}
