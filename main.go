package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	byExt := map[string][]string{}
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		ext := strings.ToLower(filepath.Ext(p))
		if ext == "" {
			ext = "(no extension)"
		}
		byExt[ext] = append(byExt[ext], p)
		return nil
	})
	fmt.Println("Suggested organization:")
	for ext, files := range byExt {
		folder := strings.TrimPrefix(ext, ".")
		fmt.Printf("\n%s/ (%d files)\n", folder, len(files))
		for _, f := range files {
			fmt.Printf("  -> %s/%s\n", folder, filepath.Base(f))
		}
	}
}
