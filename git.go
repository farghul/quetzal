package main

import (
	"os/exec"
	"strings"
)

// Switch to the main branch, and pull any changes
func prepare() {
	execute("-v", "git", "checkout", "main")
	execute("-v", "git", "pull")
}

// Create an update branch if necessary
func checkout() {
	if exists(branch, ticket) {
		execute("-v", "git", "checkout", branch+ticket)
	} else {
		execute("-v", "git", "checkout", "-b", branch+ticket)
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

// Add and commit the updates
func commit() {
	execute("-v", "git", "add", ".")
	execute("-v", "git", "commit", "-m", ticket+" install "+plugin)
}

// Tag the version so Satis can package it
func tags() {
	execute("-v", "git", "tag", "v"+satis.Version)
	execute("-v", "git", "push", "origin", "--tags")
}

// Push modified content to the git repository
func push() {
	if cherry {
		execute("-v", "git", "push", "--set-upstream", "origin", branch+ticket)
	} else {
		execute("-v", "git", "push")
	}
}
