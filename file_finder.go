package main

import (
	"fmt"
	"os"
)

func main() {

	arg := os.Args[1:]

	dirFile, err := os.OpenFile("dirs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Something wrong happend")
		return
	}
	defer func(dirFile *os.File) {
		err := dirFile.Close()
		if err != nil {

		}
	}(dirFile)

	if len(arg) < 1 {
		fmt.Println("No arguments")
		return
	}

	for i := 0; i < len(arg); i++ {
		directory := arg[i]
		files, _ := os.ReadDir(directory)
		dirFile.WriteString(directory + "               \n")
		for x := range files {
			fd, _ := files[x].Info()
			if fd.IsDir() {
				dirFile.WriteString("              " + fd.Name() + "\n")
			}
		}
	}

}
