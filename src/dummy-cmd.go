package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	input := ""

	for input != "exit" {
		fmt.Println("Waiting for command...")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		fmt.Println("Command Entered:", input)
	}
}
