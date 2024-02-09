package main

import (
	"fmt"
	"os"
)

func main() {

	file, error := os.OpenFile("empty_second.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if error != nil {
		fmt.Println("Something was wrong")
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	_, err := file.WriteString("Hello, My Good Old Friend\n")
	if err != nil {
		fmt.Printf("There was an error writing to a file named: %s\n", file.Name())
		return
	}

	fmt.Printf("Data was successfully writtent into file named: %s\n", file.Name())

}
