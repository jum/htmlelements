/*
Package htmlelements implements functions on parsed documents as returned by
the golang.org/x/net/html package. These are patterned after the known
JavaScript functions.
*/
package htmlelements

import (
	"strings"

	"golang.org/x/net/html"
)

// GetAttribute searches the named attribute, returning the value or if
// not found an empty string.
func GetAttribute(n *html.Node, attrName string) string {
	for _, a := range n.Attr {
		if a.Key == attrName {
			return a.Val
		}
	}
	return ""
}

// GetElementsByClassName searches the doc for all elements with the given
// class name.
func GetElementsByClassName(doc *html.Node, className string) []*html.Node {
	var nodes []*html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && hasClass(n, className) {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

// GetElementsByTagName searches the doc for all elements with the given
// tag name.
func GetElementsByTagName(doc *html.Node, tagName string) []*html.Node {
	var nodes []*html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == tagName {
			nodes = append(nodes, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return nodes
}

// GetElementByID searches for the HTML element with the given id.
func GetElementByID(doc *html.Node, id string) *html.Node {
	var node *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && GetAttribute(n, "id") == id {
			node = n
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return node
}

// hasClass splits the class attribute value on a space boundary and checks
// if the class argument appears in that list.
func hasClass(n *html.Node, class string) bool {
	for _, c := range strings.Split(GetAttribute(n, "class"), " ") {
		if c == class {
			return true
		}
	}
	return false
}

// InnerText retrieves all the text of all the elements concatenated.
func InnerText(n *html.Node) (ret string) {
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.TextNode {
			ret += n.Data
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(n)
	return
}
