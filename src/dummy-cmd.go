package main

import "fmt"

func main() {

	input := ""

	for input != "exit" {
		fmt.Println("Waiting for command...")
		fmt.Scanln(&input)
		fmt.Println("Command Entered: ", input)
	}
}
