package main

import (
	"fmt"
)

func main() {
	showHelp()
}

func showHelp() {
	fmt.Println(`
	
Welcome to cleanpath, a lightweight Go CLI app to clean your path from duplicates!
	
Usage: CLI Template [OPTIONS]
Options:
	-h, --help     Shows help page.
	
	`)
}
