//go:build windows

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Arthur1/windows-version-go/windows"
)

func main() {
	wv, err := windows.GetWinVer()
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
	fmt.Printf("OSName: %s\n", wv.OSName)
	fmt.Printf("Version: %s\n", wv.Version)
	fmt.Printf("Release: %s\n", wv.Release)
}
