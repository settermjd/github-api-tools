package main

import (
	"fmt"
	"flag"
	"log"
	"os"
	"encoding/json"
	"github.com/settermjd/githubapi"
)

func main() {
	var prId int
	var username string
	var password string

	flag.IntVar(&prId, "prId", 0, "The PR's id")
	flag.StringVar(&username, "username", "", "Your API username")
	flag.StringVar(&password, "password", "", "Your API password")
	flag.Parse()

	serverUri := "https://api.github.com/repos/owncloud/documentation/pulls/%d/commits"

	req := githubapi.GitRequest{ serverUri, githubapi.RequestCredentials{ username, password }}

	var commitList []githubapi.Commit

	unmarshallErr := json.Unmarshal(req.MakeRequest(prId), &commitList)

	if unmarshallErr != nil {
		log.Print(unmarshallErr)
		os.Exit(1)
	}

	commits := githubapi.Commits{ commitList }

	fmt.Println(commits.GetCommitsAsList())
}

