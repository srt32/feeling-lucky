package main

import (
	"fmt"
	"os"

	"github.com/google/go-github/github"
	"github.com/srt32/feeling-lucky/client"
	"github.com/srt32/feeling-lucky/config"
	"github.com/toqueteos/webbrowser"
)

func main() {
	//if len(os.Args) != 3 {
	//	fmt.Println("Please provide an org and repo")
	//	os.Exit(1)
	//}

	userConfig := config.UserConfig{}

	apiClient, err := client.NewClient(userConfig)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	query := os.Args[1]

	codeSearchResult, _, err := apiClient.Search.Code(
		query+"+user:thoughtbot", &github.SearchOptions{},
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//for _, result := range codeSearchResult.CodeResults {
	//	fmt.Println(result)
	//}
	fmt.Println(codeSearchResult.CodeResults[0])

	webbrowser.Open(*codeSearchResult.CodeResults[0].HTMLURL)
}
