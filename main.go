package main

import (
	"fmt"
	fbody "github.com/yashLadha/go-moji/bodyPrompt"
	ftype "github.com/yashLadha/go-moji/commitTypePrompt"
	fissue "github.com/yashLadha/go-moji/fetchissue"
	fmoji "github.com/yashLadha/go-moji/fetchmoji"
	ft "github.com/yashLadha/go-moji/fetchtitle"
	issuePrompt "github.com/yashLadha/go-moji/issueprompt"
	models "github.com/yashLadha/go-moji/models"
	"os"
	"os/exec"
)

func exitWrapper() {
	os.Exit(-1)
}

func main() {
	var commitObject models.CommitBody
	commitObject.JiraIssue = ""
	commitObject.GithubIssue = ""

	response, err := fmoji.FetchMoji()
	if err != nil {
		fmt.Printf("something unexpected occured %v\n", err)
		exitWrapper()
	}
	// fmt.Println(response)
	commitType, err := ftype.CommitTypePrompt(response)
	if err != nil {
		exitWrapper()
	}
	commitObject.Type = commitType

	titleString, err := ft.FetchTitle()
	if err != nil {
		exitWrapper()
	}

	commitObject.Title = titleString
	issueOption, err := fissue.FetchIssue()
	if err != nil {
		exitWrapper()
	}

	// Populate prompt according to issue selection
	switch issueOption {
	case "Github":
		issueString, err := issuePrompt.GithubIssuePrompt()
		if err != nil {
			exitWrapper()
		}
		commitObject.GithubIssue = issueString
		break
	case "Jira":
		issueString, err := issuePrompt.JiraIssuePrompt()
		if err != nil {
			exitWrapper()
		}
		commitObject.JiraIssue = issueString
		break
	case "NIL":
		commitObject.NoIssue = true
		break
	default:
		fmt.Println("No valid option selected")
		exitWrapper()
	}

	bodyStr, err := fbody.BodyPrompt()
	if err != nil {
		exitWrapper()
	}
	commitObject.CommitDefinition = bodyStr

	var commitStr = commitObject.Type.Code + " " + commitObject.Title + "\n"
	if commitObject.NoIssue {
		commitStr += commitObject.CommitDefinition
	} else if commitObject.JiraIssue != "" {
		commitStr += "Ref: #" + commitObject.JiraIssue + " \n" + commitObject.CommitDefinition
	} else if commitObject.GithubIssue != "" {
		commitStr += "Ref: #" + commitObject.GithubIssue + " \n" + commitObject.CommitDefinition
	}
	// fmt.Println(commitStr)
	cmd := exec.Command("git", "commit", "-m", commitStr)
	if err := cmd.Run(); err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		exitWrapper()
	}
	os.Exit(0)
}
