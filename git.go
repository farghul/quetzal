package main

import (
	"os/exec"
	"strings"
)

// Switch to the main branch, and pull any changes
func prepare() {
	execute("git", []string{"checkout", "main"}, ExecOptions{Stream: true})
	execute("git", []string{"pull"}, ExecOptions{Stream: true})
}

// Create an update branch if necessary
func checkout() {
	if exists(branch, ticket) {
		execute("git", []string{"checkout", branch + ticket}, ExecOptions{Stream: true})
	} else {
		execute("git", []string{"checkout", "-b", branch + ticket}, ExecOptions{Stream: true})
		cherry = true
	}
}

// Check to see if the current release branch already exists locally
func exists(prefix, tag string) bool {
	found := false
	b, _ := exec.Command("git", "branch").Output()
	if strings.Contains(string(b), prefix+tag) {
		found = true
	}
	return found
}

// Add and commit the update
func commit() {
	execute("git", []string{"add", "."}, ExecOptions{Stream: true})
	execute("git", []string{"commit", "-m", ticket, "-m", "Install " + plugin}, ExecOptions{Stream: true})
}

// Tag the version so Satis can package it
func tags() {
	execute("git", []string{"tag", "v" + satis.Version}, ExecOptions{Stream: true})
	execute("git", []string{"push", "origin", "--tags"}, ExecOptions{Stream: true})

}

// Push modified content to the git repository
func push() {
	if cherry {
		execute("git", []string{"push", "--set-upstream", "origin", branch + ticket}, ExecOptions{Stream: true})
	} else {
		execute("git", []string{"push"}, ExecOptions{Stream: true})
	}
}
