package commitTypePrompt

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/yashLadha/go-moji/models"
	"strings"
)

func CommitTypePrompt(emojiList *models.GitMojis) (models.EmojiInfo, error) {

	template := promptui.SelectTemplates{
		Label:    "{{ .Description }}?",
		Active:   "➜ {{ .Emoji }} {{ .Description }}",
		Inactive: "{{ .Emoji }} {{ .Description }}",
		Selected: "Commit Type {{ .Description }}",
	}

	searcher := func(input string, index int) bool {
		emoji := emojiList.Emoji[index]
		name := strings.Replace(strings.ToLower(emoji.Description), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	emojiPrompt := promptui.Select{
		Label:     "Select type of commit",
		Items:     emojiList.Emoji,
		Size:      10,
		Templates: &template,
		Searcher:  searcher,
	}

	idx, _, err := emojiPrompt.Run()
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return models.EmojiInfo{}, err
	}

	return emojiList.Emoji[idx], nil
}
