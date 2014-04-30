# Feedfinder

[![Build Status](https://travis-ci.org/pilu/feedfinder.png?branch=master)](https://travis-ci.org/pilu/feedfinder)

Feedfinder is a [Go](http://golang.org/) library that you can use to discover feed links on web pages.

```go
package main

import (
	"fmt"
	"log"

	"github.com/pilu/feedfinder"
)

func main() {
	links, err := feedfinder.Find("http://blog.golang.org")
	if err != nil {
		log.Fatal(err)
	}

	for _, link := range links {
		fmt.Printf("%s (%s): %s\n", link.Title, link.Type, link.URL.String())
	}
}
```

## Author

* [Andrea Franz](http://gravityblast.com)
