package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/google/go-github/github"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/oauth2"
	"gopkg.in/yaml.v3"
)

type Parameters struct {
	Owner       string `yaml:"owner"`
	Website     string `yaml:"website"`
	Repository  string `yaml:"repository"`
	InitTitle   string `yaml:"initTitle"`
	EndTitle    string `yaml:"endTitle"`
	Title       string
	Body        string   `yaml:"body"`
	DefaultPath string   `yaml:"defaultPath"`
	FileTarget  string   `yaml:"fileTarget"`
	Labels      []string `yaml:"labels"`
}

func createIssue(client *github.Client, param *Parameters) (*github.Issue, error) {

	fullLink := param.Website + param.FileTarget
	param.Body = strings.Replace(param.Body, "FILE_TARGET", param.FileTarget, -1)
	param.Body = strings.Replace(param.Body, "HYPERLINK", fullLink, -1)

	param.Title = param.InitTitle + "`" + param.FileTarget + "`" + param.EndTitle
	title := param.Title
	body := param.Body
	labels := param.Labels
	opt := &github.IssueRequest{Title: &title, Body: &body, Labels: &labels}
	issue, _, err := client.Issues.Create(context.Background(), param.Owner, param.Repository, opt)

	return issue, err
}

func main() {

	var yamlFile string
	fmt.Print("Enter yaml file: ")
	fmt.Scanf("%s", &yamlFile)

	// Read and Unmarshal yaml file
	absPath, _ := filepath.Abs(yamlFile)
	yfile, err := os.ReadFile(absPath)
	if err != nil {
		log.Fatal(err)
	}

	var params Parameters
	if err := yaml.Unmarshal(yfile, &params); err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}

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

	fmt.Printf("--- Create issue in %s/%s \n", params.Owner, params.Repository)
	createIssue(client, &params)

	// print the fields to the console
	fmt.Printf("fileTarget: %s\n", params.FileTarget)
	fmt.Printf("Issue Title: %s\n", params.Title)
	fmt.Printf("Issue Labels: %s\n", params.Labels)
	fmt.Println("Issue Body:")
	fmt.Printf("%s\n", params.Body)
	fmt.Println("----------------")
	//
}
