package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
)

func main() {
	pageURL := "https://www.betaarchive.com/database/sitemap.php"

	AllLinks := ScrapPageURLs(pageURL)
	OriginalReleaseLinks := CheckRelease(AllLinks)
	AbandonWareLinks := CheckAbandon(OriginalReleaseLinks)
	fmt.Println(AbandonWareLinks)
	fmt.Println("------------Finished Checking Links------------")
	SliceToFile(AbandonWareLinks)
}

func ScrapPageURLs(url string) []string {
	var Links []string
	c := colly.NewCollector()
	// fmt.Println("......Collector created......")

	c.OnHTML("url > loc", func(e *colly.HTMLElement) {
		// printing selected link
		// fmt.Println(e.DOM.Text())
		Links = append(Links, e.DOM.Text())

	})
	c.Visit(url)

	return Links
}

func SliceToFile(slice []string) {
	file, err := os.Create("Links.txt")
	if err != nil {
		fmt.Println("Error creating file")
	}
	for _, url := range slice {
		file.Write([]byte(url + "\n"))
	}
}
