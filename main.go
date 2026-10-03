package main

import (
	"fmt"
	"os"

	"github.com/JUSTINS88/githubRepoAnalyzer/ghclient"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("How to use: githubRepoAnalyzer <username>")
		return
	}

	fmt.Println("Welcome to Github Repo Analyzer!")

	usernames := os.Args[1:]

	for _, user := range usernames {
		repos, err := ghclient.FetchRepos(user)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("This is repos for %s: \n", user)
			for i, data := range repos {
				lang := data.Language
				if lang == "" {
					lang = "Unknown"
				}
				fmt.Printf("%d (stars: %d) %s Language: %s\n", i+1, data.StargazersCount, data.Name, lang)
			}
		}
		fmt.Println("-----------")
	}
}
