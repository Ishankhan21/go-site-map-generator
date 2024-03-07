package linkparser

import (
	"fmt"
	"io"
	"log"
	"strings"

	"golang.org/x/net/html"
)

type Link struct {
	Href string
	Text string
}

func Parse(reader io.Reader) ([]Link, error) {
	doc, err := html.Parse(reader)
	if err != nil {
		log.Fatal(err)
	}

	l := linkNodes(doc)
	var links []Link
	for _, n := range l {
		var link Link
		fmt.Println("Doc +++", n.Type, n.Data, n.DataAtom, n.Namespace, n.Attr)
		for _, arr := range n.Attr {
			if arr.Key == "href" {
				link.Href = arr.Val
				break
			}
		}
		link.Text = text(n)
		links = append(links, link)
	}

	fmt.Println("Links +++++++", links)
	return links, nil
}

func text(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}
	if n.Type != html.ElementNode {
		return ""
	}
	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += text(c)
	}
	return strings.Join(strings.Fields(ret), " ")
}

func linkNodes(doc *html.Node) []*html.Node {
	if doc.Type == html.ElementNode && doc.Data == "a" {
		return []*html.Node{doc}
	}
	c := doc.FirstChild
	var links []*html.Node
	for c = doc.FirstChild; c != nil; c = c.NextSibling {
		links = append(links, linkNodes(c)...)
	}
	return links
}
