package calculations

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(scanner bufio.Scanner) string {
	scanner.Scan()
	if strings.ToLower(scanner.Text()) == "exit" {
		fmt.Println("Exiting")
		os.Exit(0)
	}
	return scanner.Text()
}

func handleInputError(err error) {
	fmt.Printf("Invalid input: %s\n", err)
	fmt.Println("Try again")
}
