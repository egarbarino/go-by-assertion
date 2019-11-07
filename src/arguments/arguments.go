package main

import (
	"fmt"
	"os"
)

func main() {
	var arguments, command string
	for i := 0; i < len(os.Args); i++ {
		if i == 0 {
			command = os.Args[0]
		} else if i == 1 {
			arguments += os.Args[i]
		} else {
			arguments += " " + os.Args[i]
		}
	}
	// go run arguments.go arg1 arg2 arg3

	fmt.Println("  command: ", command)
	// command:  C:\Users\ERNEST~1\AppData\Local\Temp\go-build950394370\b001\exe\arguments.exe

	fmt.Println("arguments: ", arguments)
	// arguments:  arg1 arg2 arg3
}
