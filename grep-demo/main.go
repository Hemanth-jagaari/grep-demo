package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println("In main function")
	fmt.Printf("Type of cmd args is %T\n", args)
	fmt.Printf("cmd line arguments are %v\n", args)

	wdir, err := os.Getwd()

	if err != nil {
		panic(err)
	}
	fmt.Printf("working dir :%v\n", wdir)
}
