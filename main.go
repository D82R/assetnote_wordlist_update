package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"

	"github.com/d82r/assetnote_wordlist_update/links"
)

var automatedOutput string
var date string
var output string
var techOutput string

func main() {
	flag.StringVar(&date, "d", "", "assetnote release date (Ex: 2022_03_28)")
	flag.Parse()

	homeDirectory := os.Getenv("HOME")
	downloadDirectory := homeDirectory + "/assetnote"
	automatedOutput = homeDirectory + "/assetnote/automated"
	techOutput = homeDirectory + "/assetnote/technologies"

	autoList := links.CompileAutomated(date)
	techList := links.CompileTechnologies(date)

	getAutomatedLists(autoList, automatedOutput)
	getAutomatedLists(techList, techOutput)

	fmt.Printf("Updated wordlists downloaded to: %v\n", downloadDirectory)
}

func getAutomatedLists(list []string, output string) {
	for _, l := range list {
		cmd := "wget " + l + " -P" + output
		wgetCmd := exec.Command("bash", "-c", cmd)
		wgetCmd.Run()
	}
}
