package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	name, num_repos, err := githubInfo("a-kumar5")
	if err != nil {
		log.Fatalf("error: %s", err)
	}
	fmt.Printf("Name: %s\n", name)
	fmt.Printf("num_repos: %s", num_repos)
}

type Reply struct {
	Name     string
	NumRepos int `json:"public_repos"`
}

func githubInfo(login string) (string, int, error) {
	//loginUrl := fmt.Sprintf("https://api.github.com/users/%s", login)
	loginUrl := "https://api.github.com/users/" + url.PathEscape(login)
	resp, err := http.Get(loginUrl)
	if err != nil {
		return "", 0, err
	}
	if resp.StatusCode != http.StatusOK {
		return "", 0, err
	}
	var r Reply
	dec := json.NewDecoder(resp.Body)
	if err := dec.Decode(&r); err != nil {
		return "", 0, err
	}
	return r.Name, r.NumRepos, nil
}
