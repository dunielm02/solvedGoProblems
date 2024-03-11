package htmlParser

import (
	"io"
	"log"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func ParseDocument(document io.Reader) []Link {
	firstNode, err := html.Parse(document)
	if err != nil {
		log.Fatal(err)
	}

	return findLinks(firstNode)
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

func findLinks(node *html.Node) []Link {
	if node == nil {
		return []Link{}
	}
	if node.Type == html.ElementNode && node.Data == "a" {
		var address string
		for _, i := range node.Attr {
			if i.Key == "href" {
				address = i.Val
			}
		}
		return []Link{
			{
				Href: address,
				Text: extractText(node.FirstChild),
			},
		}
	}
	var ret []Link = findLinks(node.FirstChild)
	ret = append(ret, findLinks(node.NextSibling)...)

	return ret
}
