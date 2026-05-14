package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

// Test for a proper flag
func flags() string {
	var flag string

	if len(os.Args) == 1 {
		flag = "--zero"
	} else {
		flag = os.Args[1]
	}
	return flag
}

// Read the JSON files and Unmarshal the data into the appropriate Go structure
func serialize() {
	jsons := []string{config + "bitbucket.json", config + "credentials.json", config + "downloads.json", config + "jira.json", config + "sites.json"}
	for index, element := range jsons {
		data, err := os.ReadFile(element)
		inspect(err)
		switch index {
		case 0:
			json.Unmarshal(data, &bitbucket)
		case 1:
			json.Unmarshal(data, &auth)
		case 2:
			json.Unmarshal(data, &download)
		case 3:
			json.Unmarshal(data, &jira)
		case 4:
			json.Unmarshal(data, &sites)
		}
	}
}

// Compile the results of a Jira API query and save summary and key into a string slice
func compiler(element string) []string {
	var data []byte
	var err error

	data, err = api(jira.Basic + jira.ToDo)
	inspect(err)
	err = json.Unmarshal(data, &query)
	inspect(err)

	var candidate []string
	for i := range query.Issues {
		if strings.Contains(query.Issues[i].Fields.Summary, element) {
			candidate = append(candidate, query.Issues[i].Fields.Summary)
			candidate = append(candidate, query.Issues[i].Key)
		}
	}
	fmt.Println(candidate)
	return candidate
}

func api(criteria string) ([]byte, error) {
	baseURL := jira.URL + "search/jql?jql="

	fullURL := baseURL + criteria

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Authorization", "Basic "+jira.Token)
	req.Header.Set("Accept", "application/json")

	// Execute request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// Write a passed variable to a named file
func document(name string, d []byte) {
	inspect(os.WriteFile(name, d, 0644))
}

func execute(task string, args []string, opts ExecOptions) ([]byte, error) {
	cmd := exec.Command(task, args...)
	cmd.Env = append(os.Environ(), opts.Env...)
	cmd.Dir = opts.Dir

	if opts.Stream {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		return nil, cmd.Run()
	}

	return cmd.CombinedOutput()
}

// Check for errors, print the result if found
func inspect(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

// Println function for colourized text
func (c Color) Println(text string) {
	fmt.Println(string(c) + text + Reset)
}

// Printf function for colourized text
func (c Color) Printf(format string, a ...any) {
	fmt.Printf(string(c)+format+Reset, a...)
}

// Empty the contents a folder
func clearout(path string) {
	list := ls(path)
	for _, file := range list {
		sweep(path + file)
	}
}

// Remove files or directories
func sweep(cut ...string) {
	inspect(os.RemoveAll(cut[0.]))
}

// Record a list of files in a folder
func ls(folder string) []string {
	var content []string
	dir := expose(folder)

	files, err := dir.ReadDir(0)
	inspect(err)

	for _, f := range files {
		content = append(content, f.Name())
	}
	return content
}

// Open a file for reading and return an os.File variable
func expose(file string) *os.File {
	outcome, err := os.Open(file)
	inspect(err)
	return outcome
}
