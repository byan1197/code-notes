package main

import (
	"fmt"
	"os"
)

// should extract into modules later on

func createEntry(arg string) {
	return
}

func setup(arg string) {
	// 1. ask for directory
	// 2. ask for name
	return
}

func edit(arg string) {
	// 1. if no args, edit main page (kind of like a table of contents)
	// 2. if args, then edit the specific section
	return
}

func edit(arg string) {
	// 1. if no args, edit main page (kind of like a table of contents)
	// 2. if args, then edit the specific section
	return
}

func open(arg string) {
	var err error

	switch runtime.GOOS {
		case "linux":
			err = exec.Command("xdg-open", url).Start()
		case "windows":
			err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
		case "darwin":
			err = exec.Command("open", url).Start()
		default:
			err = fmt.Errorf("unsupported platform")
	}
	
	if err != nil {
		fmt.Errorf(err)
	}

}

func main() {
	
	const commandsToModule = map[string][]interface{
		"setup": setup,
		"create-entry": createEntry,
		"edit": edit,
		"open": open

	}
	args := os.Args

	fmt.Println(args)
}
