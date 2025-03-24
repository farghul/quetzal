package main

import (
	"fmt"
	"log"
	"os"
)

// Launch the program and execute the appropriate code
func main() {
	var flag string = flags()

	switch flag {
	case "-h", "--help":
		help()
	case "-r", "--run":
		serialize()
		prem = compiler("premium")
		if len(prem) > 0 {
			sift(prem)
		} else {
			journal("No Premium plugin update tickets to process.")
		}
		clearout(temp)
	case "-v", "--version":
		version()
	case "--zero":
		alert("No flag detected -")
	default:
		alert("Unknown argument(s) -")
		help()
	}
}

// Enter a record to the log file
func journal(message string) {
	file, err := os.OpenFile("/data/automation/logs/quetzal.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	inspect(err)
	log.SetOutput(file)
	log.Println(message)
}

// Print a colourized error message
func alert(message string) {
	fmt.Println("\n", bgred, message, halt, reset)
	fmt.Println(bgyellow, "Use -h for more detailed help information ")
	os.Exit(0)
}

// Provide and highlight informational messages
func tracking(message string) {
	fmt.Println(yellow)
	fmt.Println("**", reset, message, yellow, "**", reset)
}

// Print program version number
func version() {
	fmt.Println("\n", yellow+"quetzal", green+bv, reset)
}

// Print help information for using the program
func help() {
	fmt.Println(yellow, "\nUsage:", reset)
	fmt.Println("  [program] [flag]")
	fmt.Println(yellow, "\nOptions:")
	fmt.Println(green, " -h, --help", reset, "      Help Information")
	fmt.Println(green, " -r, --run", reset, "       Run Program")
	fmt.Println(green, " -v, --version", reset, "   Display Program Version")
	fmt.Println(yellow, "\nExample:", reset)
	fmt.Println("   quetzal -r")
	fmt.Println(yellow, "\nHelp:", reset)
	fmt.Println("  For more information go to:")
	fmt.Println(green, "   https://github.com/farghul/quetzal.git")
	fmt.Println(reset)
}
