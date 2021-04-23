package main

import (
	"fmt"
	"os"
)

func main() {
	var root = "c:/repos"
	var dirs = make([]string, 25)
	getDirs(root, dirs)
}
func getDirs(dir string, dirs []string) {
	f, err := os.Open("c:/repos")
	if err != nil {
		fmt.Printf("error opening directory: %s", err)
		return
	}
	fns, err := f.Readdirnames(0)
	if err != nil {
		fmt.Printf("error opening directory: %s", err)
		return
	}
	var ndirs = append(dirs, fns)
}
