package issueprompt

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"regexp"
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
	re := regexp.MustCompile(`([A-Z]+-[\d]+)`)
	validator := func(input string) error {
		matched := re.MatchString(input)
		if !matched {
			return errors.New("Wrong Issue Format")
		}
		return nil
	}
	issuePrompt := promptui.Prompt{
		Label:    "Jira Issue",
		Validate: validator,
	}

	issueTag, err := issuePrompt.Run()
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return "", err
	}
	return issueTag, nil
}
