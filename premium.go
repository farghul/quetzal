package main

import (
	"encoding/json"
	"os"
	"strings"
)

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

// Premium directs the preliminary actions to determine if the program can continue
func premium() {
	os.Chdir(repos + folder[1])
	learn()

	/*
		// This section does not work as expected, need to come back to it.
		switch folder[1] {
		case "polylang-pro":
			login(auth.Auth[1].Username, auth.Auth[1].Password, download.Polylang, sites.Polylang)
			execute("-v", "unzip", temp+folder[1]+".zip", "-d", temp)
		case "uji-countdown-premium":
			execute("-v", "curl", "-L", download.Uji, "-o", temp+folder[1]+".zip")
			execute("-v", "unzip", temp+folder[1]+".zip", "-d", temp)
		case "wp-all-export-pro":
			login(auth.Auth[2].Username, auth.Auth[2].Password, download.AllExport, sites.AllExport)
			execute("-v", "unzip", temp+folder[1]+".zip", "-d", temp)
		}
	*/

	satis.Version = number[1]

	if satis.Name+":"+satis.Version == plugin {
		quarterback()
	} else {
		alert("Plugin name does not match composer.json entry - ")
	}
}

/*
// This function is releated to unused code.
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

	execute("-v", "curl", "-L", download, "-o", temp+folder[1]+".zip")
}
*/

// Read the composer.json file and store the results in a structure
func learn() {
	current, _ := os.ReadFile("composer.json")
	err := json.Unmarshal(current, &satis)
	inspect(err)
}

// A sequential list of tasks run to complete the program
func quarterback() {
	prepare()
	inform("Switching to update branch")
	checkout()
	inform("Running update script")
	script()
	correct()
	inform("Commiting changes")
	commit()
	inform("Tagging to Satis")
	tags()
	inform("Pushing to repository")
	push()
}

// Run the update script on downloaded content
func script() {
	execute("sh", []string{"-c", "scripts/update.sh " + temp + folder[1] + "/"}, ExecOptions{Stream: true})
}

// Convert the structure back into json and overwrite the composer.json file
func correct() {
	var updated []byte
	updated, _ = json.MarshalIndent(satis, "", "    ")
	document("composer.json", updated)
}
