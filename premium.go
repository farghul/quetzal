package main

import (
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"golang.org/x/net/publicsuffix"
)

// A sequential list of tasks run to complete the program
func quarterback() {
	prepare()
	tracking("Switching to update branch")
	checkout()
	tracking("Running update script")
	script()
	correct()
	tracking("Commiting changes")
	commit()
	tracking("Tagging to Satis")
	tags()
	tracking("Pushing to repository")
	push()
}

// Premium directs the preliminary actions to determine if the program can continue
func premium() {
	os.Chdir(repos + folder[1])
	learn()

	switch folder[1] {
	case "events-calendar-pro":
		execute("-v", "curl", "-L", download.Calendar, "-o", assets+"temp/"+folder[1]+".zip")
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	case "event-tickets-plus":
		execute("-v", "curl", "-L", download.Tickets, "-o", assets+"temp/"+folder[1]+".zip")
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	case "events-virtual":
		execute("-v", "curl", "-L", download.Virtual, "-o", assets+"temp/"+folder[1]+".zip")
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	case "gravityforms":
		login(cred[0].Username, cred[0].Password, download.Gravity, site.Gravity)
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	case "polylang-pro":
		login(cred[1].Username, cred[1].Password, download.Polylang, site.Polylang)
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	case "uji-countdown-premium":
		execute("-v", "curl", "-L", download.Uji, "-o", assets+"temp/"+folder[1]+".zip")
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	case "wp-all-export-pro":
		login(cred[2].Username, cred[2].Password, download.AllExport, site.AllExport)
		execute("-v", "unzip", assets+"temp/"+folder[1]+".zip", "-d", assets+"temp/")
	}

	satis.Version, ecp.Version, evtp.Version = number[1], number[1], number[1]

	if strings.Contains(folder[1], "event") {
		if ecp.Name+":"+ecp.Version == plugin || evtp.Name+":"+evtp.Version == plugin {
			quarterback()
		}
	} else if satis.Name+":"+satis.Version == plugin {
		quarterback()
	} else {
		alert("Plugin name does not match composer.json entry - program halted")
	}
}

// Iterate through the Args array and assign plugin and ticket values
func sift(box []string) {
	for i := 0; i < len(box); i++ {

		plugin = box[i]
		i++
		ticket = box[i]

		assign()
		premium()
	}
}

// Split the supplied arguments and assign them to variables
func assign() {
	number = strings.Split(plugin, ":")
	folder = strings.Split(number[0], "/")
}

// Read the composer.json file and store the results in a structure
func learn() {
	current, _ := os.ReadFile("composer.json")
	err := json.Unmarshal(current, &satis)
	inspect(err)
	err = json.Unmarshal(current, &ecp)
	inspect(err)
	err = json.Unmarshal(current, &evtp)
	inspect(err)
}

// Process to pass credentials and download a zip file
func login(username, password, download, login string) {
	options := cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(&options)
	inspect(err)
	client := http.Client{Jar: jar}
	client.PostForm(login, url.Values{
		"password": {password},
		"username": {username},
	})

	execute("-v", "curl", "-L", download, "-o", assets+"temp/"+folder[1]+".zip")
}

// Run the update script on downloaded content
func script() {
	execute("-v", "sh", "-c", "scripts/update.sh "+assets+"temp/"+folder[1]+"/")
}

// Convert the structure back into json and overwrite the composer.json file
func correct() {
	var updated []byte
	if strings.Contains(ecp.Name, "calendar") {
		updated, _ = json.MarshalIndent(ecp, "", "    ")
	} else if strings.Contains(evtp.Name, "tickets") || strings.Contains(evtp.Name, "virtual") {
		updated, _ = json.MarshalIndent(evtp, "", "    ")
	} else {
		updated, _ = json.MarshalIndent(satis, "", "    ")
	}
	document("composer.json", updated)
}
