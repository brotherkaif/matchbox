package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter a message:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured whilst attempting to read input. Please try again.", err)
		return
	}

	fmt.Println("The message you entered was:", input)
}
