package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	// "os/exec"
	// "runtime"
)

// should extract into modules later on
func createEntry(arg string) {
	return
}

func codeNoteInit(arg string) {
	// 1. ask for directory
	// 2. ask for name
	directory := "~/Documents/code-notes"
	dirScan := "~/"
	var name string
	fmt.Printf("Enter a directory for your code diary. (Default %s)\n", directory)
	fmt.Scanln(&dirScan)
	if dirScan != "" {
		directory = dirScan
	}
	fmt.Printf("Got it. You want your notes to be stored in %s", directory)
	os.MkdirAll(directory, os.ModePerm)

	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	return
}

func edit(arg string) {
	// 1. if no args, edit main page (kind of like a table of contents)
	// 2. if args, then edit the specific section
	return
}

func open(arg string) {
	// var err error

	// switch runtime.GOOS {
	// case "linux":
	// 	err = exec.Command("xdg-open", url).Start()
	// case "windows":
	// 	err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	// case "darwin":
	// 	err = exec.Command("open", url).Start()
	// default:
	// 	err = fmt.Errorf("unsupported platform")
	// }

	// if err != nil {
	// 	fmt.Errorf(err)
	// }

}

type codeNoteConfiguration struct {
	notesDirectory string `json:"notesDirectory"`
	name           string `json:"name"`
}

func getConfig() codeNoteConfiguration {
	file, _ := ioutil.ReadFile("~/.config-helper.json")

	configContents := codeNoteConfiguration{}
	_ = json.Unmarshal([]byte(file), &configContents)

	return configContents
}

func main() {

	commandsToModule := map[string]interface{}{
		"setup":        codeNoteInit,
		"create-entry": createEntry,
		"edit":         edit,
		"open":         open,
		"help":         open}

	args := os.Args

	configuration := getConfig()

	if configuration == (codeNoteConfiguration{}) && args[1] != "setup" {
		fmt.Println("No configuration file. Please run \n\n\tcode-notes setup")
		return
	}

	fn, foundFunction := commandsToModule[args[1]].(func(string))

	if foundFunction {
		fn("asdf")
		return
	}

	fmt.Printf("%s is not a valid command command. See 'code-notes help'", args[1])
}
