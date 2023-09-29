package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) == 2 {
		file, err := os.Open(args[1])
		if err != nil {
			fmt.Printf("what i want is : %v\n", err.Error())
		}

		arr := make([]byte, 14)

		file.Read(arr)

		fmt.Printf("%s", arr)
		fmt.Println()

		file.Close()
	} else if len(args) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("File name missing")
	}
}
