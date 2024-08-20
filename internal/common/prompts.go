package common

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func MutedMessage(message string) string {
	return fmt.Sprintf("%s%s%s", colorGray, message, colorNone)
}

func PromptBool(message string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + " (y/N): ")

	var err error

	input, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Error reading input - " + err.Error())
	}

	// Remove the newline character and convert to lower case
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	if input != "y" && input != "yes" {
		os.Exit(0)
	}

	return nil
}

func PromptEnter(message string) error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message)

	_, err := reader.ReadString('\n')
	if err != nil {
		return errors.New("Error reading input - " + err.Error())
	}

	// Clear the line by overwriting it with spaces and then returning to the start of the line
	clearLine := "\r" + strings.Repeat(" ", len(message)) + "\r"
	fmt.Print(clearLine)

	return nil
}
