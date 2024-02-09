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

	data := []byte("Hello World!\n")
	err = os.WriteFile("empty_second.txt", data, 0644)

	if err != nil {
		fmt.Println("Error writing to a file: ", err)
		return
	}

	fmt.Println("Data written successfully to empty_second")

	var names []string

	for _, file := range files {
		info, err := file.Info() // Get FileInfo for each file/directory
		if err != nil {
			fmt.Println("Error retrieving file info:", err)
			continue
		}

		// Check if the file size is 0 (file is empty)
		if info.Size() == 0 {
			names = append(names, info.Name())
		}
	}

	fmt.Println(names)

}
