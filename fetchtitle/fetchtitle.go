package fetchtitle

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
)

// FetchTitle Creates a commit title prompt for fetching input from user
func FetchTitle() (commitTitle string, err error) {
	validator := func(input string) error {
		if len(input) >= 50 {
			return errors.New("Commit Title length execeeded (Max len: 50)")
		}
		return nil
	}

	titleWritingPrompt := promptui.Prompt{
		Label:    "Enter commit title",
		Validate: validator,
	}

	result, err := titleWritingPrompt.Run()
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return "", err
	}
	return result, nil
}
