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

	input, err := reader.ReadString('\n')
	if err != nil {
		// return a specific string when a user hits ctrl+c
		if err.Error() == "EOF" {
			return errors.New("user cancelled the operation")
		}

		return errors.New("error reading input - " + err.Error())
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
		// return a specific string when a user hits ctrl+c
		if err.Error() == "EOF" {
			return errors.New("user cancelled the operation")
		}

		return errors.New("error reading input - " + err.Error())
	}

	return nil
}
