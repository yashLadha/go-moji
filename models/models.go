package models

type GitMojis struct {
	Emoji []EmojiInfo `json:"gitmojis"`
}

type EmojiInfo struct {
	Emoji       string `json:"emoji"`
	Entity      string `json:"entity"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

// CommitBody structure of go-moji commit
type CommitBody struct {
	Title            string
	Type             EmojiInfo
	CommitDefinition string
	GithubIssue      string
	JiraIssue        string
	NoIssue          bool
}
