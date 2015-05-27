package main

import (
	"fmt"

	"github.com/moovweb/gokogiri"
)

func main() {

	html := "<div><h1>Hello, World!</h1></div>"

	parser, _ := gokogiri.ParseHtml([]byte(html))
	defer parser.Free()

	divs, _ := parser.Search("//div")

	for _, div := range divs {
		content := div.Content()
		fmt.Printf("type: %T value: %v\n", content, content)
	}
}
