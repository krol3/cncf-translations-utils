package main

import (
	"context"
	"fmt"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
)

// Fetch all the public organizations' membership of a user.
func fetchOrganizations(client *github.Client, username string) ([]*github.Organization, error) {
	orgs, _, err := client.Organizations.List(context.Background(), username, nil)
	return orgs, err
}

func fetchRepositories(client *github.Client, username string) ([]*github.Repository, error) {
	// list public repositories for org "github"
	opt := &github.RepositoryListOptions{Type: "owner"}
	repos, _, err := client.Repositories.List(context.Background(), username, opt)

	return repos, err
}

func createIssue(client *github.Client, owner, repo string) (*github.Issue, error) {
	title := "Test title"
	body := "Test body"
	labels := []string{"test-bug"}
	opt := &github.IssueRequest{Title: &title, Body: &body, Labels: &labels}
	issue, _, err := client.Issues.Create(context.Background(), owner, repo, opt)

	return issue, err
}

func main() {
	fmt.Print("GitHub Token: ")
	byteToken, _ := terminal.ReadPassword(int(syscall.Stdin))
	println()
	token := string(byteToken)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// list all repositories for the authenticated user
	repos, _, err := client.Repositories.List(ctx, "", nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	for i, rep := range repos {
		fmt.Printf("%v. %v\n", i+1, rep.GetFullName())
	}

	var username string
	fmt.Print("Enter GitHub username: ")
	fmt.Scanf("%s", &username)

	organizations, err := fetchOrganizations(client, username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("--- Fetch all the public organizations' membership of a user [%s] \n", username)
	for i, organization := range organizations {
		fmt.Printf("%v. %v\n", i+1, organization.GetLogin())
	}

	repos, err = fetchRepositories(client, username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Printf("--- Fetch all the private repositories' membership of a user [%s] \n", username)
	for i, rep := range repos {
		fmt.Printf("%v. %v\n", i+1, rep.GetFullName())
	}

	fmt.Printf("--- Create issue in [%s]/[%s] \n", username, username)
	createIssue(client, "kcloudn", "test-private")
}
