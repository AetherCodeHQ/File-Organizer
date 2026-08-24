
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
	var n int
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		n++
		fmt.Println(p)
		return nil
	})
	fmt.Printf("%d file(s)\n", n)
}
