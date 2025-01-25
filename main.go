package main

// Launch the program and execute the appropriate code
func main() {
	var flag string = flags()
	clearout(assets + "temp/")

	switch flag {
	case "-h", "--help":
		help()
	case "-v", "--version":
		version()
	case "--zero":
		serialize()
		prem = compiler("premium")
		if len(prem) > 0 {
			sift(prem)
		} else {
			journal("No Premium plugin update tickets to process.")
		}
	default:
		alert("Unknown argument(s) -")
		help()
	}
}
