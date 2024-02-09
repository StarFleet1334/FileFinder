package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a directory")
		return
	}

	files, err := os.ReadDir(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		info, err := file.Info() // Get FileInfo for each file/directory
		if err != nil {
			fmt.Println("Error retrieving file info:", err)
			continue
		}

		// Check if the file size is 0 (file is empty)
		if info.Size() == 0 {
			fmt.Println(file.Name()) // Print the name of the empty file
		}
	}

}
