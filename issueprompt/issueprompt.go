package issueprompt

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"strconv"
)

// GithubIssuePrompt Takes Github Issue number as input
func GithubIssuePrompt() (issuenumber string, err error) {
	validator := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid Issue Number")
		}
		return nil
	}

	issuePrompt := promptui.Prompt{
		Label:    "Github Issue Number",
		Validate: validator,
	}

	issueNumber, err := issuePrompt.Run()
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return "", err
	}
	return issueNumber, nil
}

// JiraIssuePrompt Takes Jira Issue input from user
func JiraIssuePrompt() (issuenumber string, err error) {
	issuePrompt := promptui.Prompt{
		Label: "Jira Issue",
	}

	issueTag, err := issuePrompt.Run()
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return "", err
	}
	return issueTag, nil
}
