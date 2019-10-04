// Simple utility to convert a file into a Go byte array

// Clint Caywood

// http://github.com/cratonica/2goarray
package main

import (
	"fmt"
	"io"
	"os"
)

const (
	NAME         = "2goarray"
	VERSION      = "0.1.0"
	URL          = "http://github.com/cratonica/2goarray"
	GENERATED_BY = "// File generated by " + NAME + " v" + VERSION + " (" + URL + ")"
)

func main() {
	var arrayName, packageName string
	if len(os.Args) == 2 {
		arrayName = os.Args[1]
	} else if len(os.Args) == 3 {
		arrayName = os.Args[1]
		packageName = os.Args[2]
	} else {
		fmt.Print(NAME + " v" + VERSION + "\n\n")
		fmt.Println("Usage: " + NAME + " array_name (optional:package_name)")
		return
	}

	if isTerminal() {
		fmt.Println("\nPlease pipe the file you wish to encode into stdin\n")
		return
	}

	fmt.Println(GENERATED_BY + "\n")
	if packageName != "" {
		fmt.Printf("package %s\n\n", packageName)
	}
	fmt.Printf("var %s []byte = []byte{", arrayName)
	buf := make([]byte, 1)
	var err error
	var totalBytes uint64
	var n int
	for n, err = os.Stdin.Read(buf); n > 0 && err == nil; {
		if totalBytes%12 == 0 {
			fmt.Printf("\n\t")
		}
		fmt.Printf("0x%02x, ", buf[0])
		totalBytes++
		n, err = os.Stdin.Read(buf)
	}
	if err != nil && err != io.EOF {
		fmt.Errorf("Error: %v", err)
	}
	fmt.Print("\n}\n\n")
}
