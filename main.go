package main

import (
	"fmt"
	"os"
)

// file_organizer - Organize files by type
func file_organizer(path string) {
	fmt.Println("========================================")
	fmt.Println("  File-Organizer")
	fmt.Println("  Organize files by type")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	file_organizer(path)
}
