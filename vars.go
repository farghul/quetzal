package main

type ExecOptions struct {
	Stream bool
	Env    []string
	Dir    string
}

// BitBucket builds a list of BitBucket tokens and API addresses
type BitBucket struct {
	Email     string `json:"email"`
	Token     string `json:"token"`
	URL       string `json:"url"`
	UUID      string `json:"uuid"`
	Reviewers struct {
		One   string `json:"one"`
		Two   string `json:"two"`
		Three string `json:"three"`
	}
}

type Authentication struct {
	Credentials []struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"credentials"`
}

// Downloads contains the direct download links
type Downloads struct {
	Polylang  string `json:"polylang"`
	AllExport string `json:"wp-all-export"`
	Spotlight string `json:"spotlight"`
	Uji       string `json:"uji-countdown"`
}

// Jira builds a list of jira tokens and API addresses
type Jira struct {
	Token string `json:"token"`
	ToDo  string `json:"todo"`
	Basic string `json:"basic"`
	URL   string `json:"url"`
}

// JQL holds the extracted data from the JQL queries
type JQL struct {
	Issues []struct {
		ID     string `json:"id"`
		Key    string `json:"key"`
		Fields struct {
			Summary string `json:"summary"`
		} `json:"fields"`
	} `json:"issues"`
}

// Satis structure captures the contents of the composer.json file for typical premium plugins
type Satis struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Type    string `json:"type"`
}

// Sites holds the login URLs to access the premium plugin accounts
type Sites struct {
	Polylang  string `json:"polylang"`
	AllExport string `json:"wp-all-export"`
	Spotlight string `json:"spotlight"`
}

type Color string

// List of constant values
const (
	Reset           = "\033[0m"
	Black    Color  = "\033[30m"
	Red      Color  = "\033[31m"
	Green    Color  = "\033[32m"
	Yellow   Color  = "\033[33m"
	Blue     Color  = "\033[34m"
	Magenta  Color  = "\033[35m"
	Cyan     Color  = "\033[36m"
	White    Color  = "\033[37m"
	BGRed    Color  = "\033[41m"
	BGYellow Color  = "\033[43m"
	bv       string = "1.0.0"
	branch   string = "update/"
	halt     string = "program halted "
	// temp     string = "/data/automation/temp/"
	// repos    string = "/data/automation/repos/"
	// config   string = "/data/automation/jsons/"
	temp   string = "/Users/bstuike/Documents/local/temp/"
	repos  string = "/Users/bstuike/Documents/bitbucket/"
	config string = "/Users/bstuike/Documents/local/resources/"
)

// List of variables
var (
	query                JQL
	jira                 Jira
	satis                Satis
	sites                Sites
	plugin               string
	ticket               string
	download             Downloads
	bitbucket            BitBucket
	auth                 Authentication
	cherry               = false
	folder, number, prem []string
)
