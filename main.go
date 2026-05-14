package main

import (
	"fmt"
	"os"
)

// Launch the program and execute the appropriate code
func main() {
	var flag string = flags()
	logo()

	switch flag {
	case "-h", "--help":
		help()
	case "-r", "--run":
		credits()
		serialize()
		prem = compiler("premium-plugin")
		if len(prem) > 0 {
			sift(prem)
		} else {
			fmt.Println("No Premium plugin update tickets to process.")
		}
		clearout(temp)
	case "-v", "--version":
	case "--zero":
		alert("No flag detected - ")
	default:
		alert("Unknown argument(s) - ")
		help()
	}
}

// Provide and highlight an informational message
func inform(message string) {
	Yellow.Printf("%s", "** ")
	fmt.Print(message)
	Yellow.Println(" **")
}

// Print a colourized error message
func alert(message string) {
	Red.Printf("\n%s", "Error: ")
	fmt.Printf("%s", message)
	BGRed.Println(halt)
	inform("Use -h to display help information")
	os.Exit(0)
}

// Print help information for using the program
func help() {
	Yellow.Println("\nUsage:")
	fmt.Println("  [program] [flag]")
	Yellow.Println("\nOperational Flags:")
	Green.Printf("%s", "  -h, --help")
	fmt.Println("		Help Information")
	Green.Printf("%s", "  -r, --run")
	fmt.Println("		Run Program")
	Green.Printf("%s", "  -v, --version")
	fmt.Println("		Display Program Version")
	Yellow.Println("\nExample:")
	fmt.Println("  Adding your path to file if necessary, run:")
	Green.Printf("%s", "    quetzal -r")
	Yellow.Println("\nHelp:")
	fmt.Println("  For more information go to:")
	Green.Println("    https://github.com/farghul/quetzal.git")
}

// Print the program logo and build version
func logo() {
	Blue.Println("▗▄▄▄▖ ▗▖ ▗▖▗▄▄▄▖▗▄▄▄▖▗▄▄▄▄▖ ▗▄▖ ▗▖   ")
	Blue.Println("▐▌ ▐▌ ▐▌ ▐▌▐▌     █     ▗▞▘▐▌ ▐▌▐▌   ")
	Blue.Println("▐▌ ▐▌ ▐▌ ▐▌▐▛▀▀▘  █   ▗▞▘  ▐▛▀▜▌▐▌   ")
	Blue.Println("▐▙▄▟▙▖▝▚▄▞▘▐▙▄▄▖  █  ▐▙▄▄▄▖▐▌ ▐▌▐▙▄▄▖")
	Blue.Println(bv)
}

// Print program function and creator credit
func credits() {
	fmt.Println("\nA ticket creation tool for Jira")
	fmt.Println("Created by Byron Stuike")
}
