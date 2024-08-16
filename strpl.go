package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage : <program> <old-string> <new-string> <file.txt>")
		return
	}

	var oldstring string = os.Args[1]
	var newstring string = os.Args[2]
	var name string = os.Args[3]

	var file, fileerror = os.OpenFile(name, os.O_RDWR, 0660)
	if fileerror != nil {
		fmt.Println(fileerror)
	}

	var stat, staterror = file.Stat()
	if staterror != nil {
		fmt.Println(staterror)
	}

	var size = stat.Size()
	var bytes = make([]byte, size)
	file.Read(bytes)
	var content = string(bytes)
	var newcontent string = strings.ReplaceAll(string(content), oldstring, newstring)

	// DEBUG MODE
	var debug bool = false // ENABLE DEBUG MODE
	if debug {
		fmt.Println("CURRENT CONTENT")
		fmt.Println(string(content))
		fmt.Println("NEW CONTENT")
		fmt.Println(string(newcontent))
	}

	if content == newcontent {
		fmt.Println("No changes were made.")
	} else {
		file.Truncate(0)
		file.Seek(0, 0)
		file.WriteString(newcontent)
		file.Sync()
	}

}
