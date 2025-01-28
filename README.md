# Quetzal

Quetzal is a WordPress premium (subscription based) plugin update install tool. It queries a Jira API to find and simplify the process of updating WordPress plugins, while still tracking them via Jira tickets. Meant for an environment where strict version control is needed. Named after the Quetzal bird whose image graces the currency of Guatemala.

![Bird](quetzal.webp)

## Prerequisites

Googles' [Go language](https://go.dev) installed to enable building executables from source code.

Login information to download update packages.

An selection of `json` files to enable authorized querying, and everything needed to aquire the Premium plugin update files (see `jsons` folder for reference).

## Function

Quetzal searches the targeted Jira API for tickets marked as **"New"** (aka ToDo), with a summary containing the `premium-plugin` vendor tag . It then gathers the qualifying candidates, downloads and extracts the update files, and runs an `update.sh` script for each individual plugin. Finally the new version is tagged and pushed to a designated update branch.

## Build

Before building the application, change the value of the `assets` and `repos` constants to reflect your environment:

``` go
assets string = "/data/scripts/automation/assets/"
repos string = "/data/scripts/automation/bitbucket/"
```

Then, from the root folder containing `main.go`, use the command that matches your environment:

### Windows & Mac:

``` zsh
go build -o [name] .
```

### Linux:

``` zsh
GOOS=linux GOARCH=amd64 go build -o [name] .
```

## Run

``` zsh
[program] [flag]
```

## Options

``` zsh
-h, --help        Help information
-r, --run         Run program
-v, --version     Display program version
```

## Example

``` zsh
quetzal -r
```

## License

Code is distributed under [The Unlicense](https://github.com/farghul/quetzal/blob/main/LICENSE.md) and is part of the Public Domain.
