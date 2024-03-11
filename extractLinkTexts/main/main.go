package main

import (
	"flag"
	"fmt"
	"htmlParser"
	"log"
	"net/http"
)

func main() {
	page := flag.String("a", "https://www.hivac.io/es", "Change the address to find links")
	flag.Parse()

	req, err := http.NewRequest("GET", *page, nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	links := htmlParser.ParseDocument(res.Body)

	for _, i := range links {
		fmt.Printf("{\n\t%s\n\t%s\n}\n", i.Href, i.Text)
	}
}
