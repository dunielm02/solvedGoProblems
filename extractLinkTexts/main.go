package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type Link struct {
	href string
	text string
}

var links []Link = make([]Link, 0)

func ParseDocument(document io.Reader) *html.Node {
	node, err := html.Parse(document)
	if err != nil {
		log.Fatal(err)
	}

	return node
}

func extractText(node *html.Node) string {
	if node == nil {
		return ""
	}
	if node.Type == html.TextNode {
		return node.Data + extractText(node.FirstChild) + extractText(node.NextSibling)
	}
	return extractText(node.FirstChild) + extractText(node.NextSibling)
}

func findLinks(node *html.Node) {
	if node == nil {
		return
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		var address string
		for _, i := range node.Attr {
			if i.Key == "href" {
				address = i.Val
			}
		}
		links = append(links, Link{
			href: address,
			text: extractText(node),
		})
	}
	findLinks(node.NextSibling)
	findLinks(node.FirstChild)
}

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

	firstNode := ParseDocument(res.Body)

	findLinks(firstNode)

	for _, i := range links {
		fmt.Printf("{\n\t%s\n\t%s\n}\n", i.href, i.text)
	}
}
