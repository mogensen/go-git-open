package main

import (
	"fmt"
	"log"

	"github.com/bacongobbler/browser"
)

func openBrowser(url string) {
	if Print {
		fmt.Println(url)
		return
	}

	if err := browser.Open(url); err != nil {
		log.Fatal(err)
	}
}
