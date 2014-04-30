package main

import (
	"fmt"
	"os"

	"github.com/pilu/feedfinder"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage:\n  %s URL\n", os.Args[0])
		os.Exit(1)
	}

	url := os.Args[1]
	links, err := feedfinder.Find(url)
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("FEEDS FOR %s\n\n", url)
	for i, link := range links {
		fmt.Printf("%d) %s (%s)\n\n\t%s\n\n", i+1, link.Title, link.Type, link.URL.String())
	}
}
