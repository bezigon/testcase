package main

import (
	"fmt"

	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xml"
)

func getContent(node xml.Node) (string, error) {
	var content string
	switch t := node.(type) {
	case *xml.TextNode:
		content = node.Content()
	case *xml.ElementNode:
		child := node.FirstChild()
		if child == nil {
			return "", fmt.Errorf("child not found")
		}
		content = child.Content()
	default:
		return "", fmt.Errorf("unexpected type %T\n", t)
	}
	return content, nil
}

func main() {

	html := "<div><h1>Hello, World!</h1></div>"

	parser, err := gokogiri.ParseHtml([]byte(html))
	if err != nil {
		panic(err)
	}
	defer parser.Free()

	divs, _ := parser.Search("//div")

	for _, div := range divs {
		content, _ := getContent(div)
		fmt.Printf("type: %T value: %v\n", content, content)
	}
}
