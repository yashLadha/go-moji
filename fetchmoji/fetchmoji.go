package fetchmoji

import (
	"encoding/json"
	"fmt"
	"github.com/yashLadha/go-moji/models"
	"io/ioutil"
	"net/http"
)

// FetchMoji Fetched emoji data from gitmoji
func FetchMoji() (*models.GitMojis, error) {
	res, err := http.Get("https://raw.githubusercontent.com/carloscuesta/gitmoji/master/src/data/gitmojis.json")
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return new(models.GitMojis), err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Unexpected error occured %v\n", err)
		return new(models.GitMojis), err
	}
	var mojiModel = new(models.GitMojis)
	marshallErr := json.Unmarshal(body, &mojiModel)
	if marshallErr != nil {
		fmt.Printf("Unexpected error occured %v\n", marshallErr)
		return new(models.GitMojis), err
	}
	return mojiModel, nil
}
