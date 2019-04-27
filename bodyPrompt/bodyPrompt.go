package bodyPrompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

func BodyPrompt() (string, error) {
	bodyDefPrompt := promptui.Prompt{
		Label: "Enter commit body ",
	}

	result, err := bodyDefPrompt.Run()
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return "", err
	}
	return result, nil
}
