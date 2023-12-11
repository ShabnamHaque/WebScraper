package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	url := "http://quotes.toscrape.com/page/1/"

	//The callback will be executed whenever an HTML
	//element with the class "quote" is encountered on the page
	c.OnHTML(".quote", func(e *colly.HTMLElement) {
		// Extract data from HTML elements
		quote := e.ChildText("span.text")
		/* is assigned the text content of the <span class="text">
		element within the "quote" element.*/
		author := e.ChildText("small.author")
		tags := e.ChildText("a.tag")

		// Clean up the extracted data - remove leading & trailing whitespace
		quote = strings.TrimSpace(quote)
		author = strings.TrimSpace(author)
		tags = strings.TrimSpace(tags)

		// Print the scraped data
		fmt.Printf("Quote: %s\nAuthor: %s\nTags: %s\n ------ \n", quote, author, tags)
	})

	// Visit the URL and start scraping
	err := c.Visit(url)
	if err != nil {
		log.Fatal(err)
	}
}
