package ghclient

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Repo struct {
	Name            string `json:"name"`
	StargazersCount int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
	Language        string `json:"language"`
}

func FetchRepos(username string) ([]Repo, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/repos", username)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	var data []Repo
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}
