package htmlparser

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

type Link struct {
	Text string
	Href string
}

func Parser(r io.Reader) (links []Link) {
	doc, _ := html.Parse(r)
	nodes := buildNodes(doc)

	for _, node := range nodes {
		links = append(links, buildLinks(node))
	}

	for _, link := range links {
		fmt.Println(link.Href, link.Text)
	}

	return
}

func buildLinks(node *html.Node) (res Link) {
	for _, attr := range node.Attr {
		if attr.Key == "href" {
			res.Href = attr.Val
			break
		}
	}

	res.Text = getText(node)

	return
}

func buildNodes(node *html.Node) (res []*html.Node) {
	if node.Type == html.ElementNode && node.Data == "a" {
		return []*html.Node{node}
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		res = append(res, buildNodes(c)...)
	}

	return
}

func getText(node *html.Node) string {
	if node.Type == html.TextNode {
		return node.Data
	}

	if node.Type != html.ElementNode {
		return ""
	}

	var res string
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		res += getText(c)
	}

	return strings.Join(strings.Fields(res), " ")
}

//temp func.
func dfs(node *html.Node, padding string) {
	msg := node.Data
	if node.Type == html.ElementNode {
		msg = padding + "<" + msg + ">"
	}

	fmt.Println(padding, msg)

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		dfs(c, " ")
	}
}
