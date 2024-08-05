package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func PromptBool(message string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + " (y/N): ")

	var err error

	input, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Error reading file - " + err.Error())
	}

	// Remove the newline character and convert to lower case
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	if input != "y" && input != "yes" {
		os.Exit(0)
	}

	return nil
}
