package fetchissue

import (
	"fmt"
	"github.com/manifoldco/promptui"
)

// FetchIssue fetches the issue type from user
func FetchIssue() (result string, err error) {
	issueSelectionPrompt := promptui.Select{
		Label: "Select Issue Type",
		Items: []string{"Github", "Jira", "NIL"},
	}
	_, option, err := issueSelectionPrompt.Run()
	if err != nil {
		fmt.Printf("Something unexpected occured %v\n", err)
		return "", err
	}
	return option, nil
}
