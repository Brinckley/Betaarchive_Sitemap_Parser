package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

func CheckOneRelease(url string) bool {
	IsOriginal := false
	c := colly.NewCollector()

	c.OnHTML("div.col-md-6 > table tr", func(e *colly.HTMLElement) {
		column1 := e.DOM.Find("td:nth-child(1)").Text()
		column2 := e.DOM.Find("td:nth-child(2)").Text()
		if strings.Contains(column1, "Original Release?") {
			if strings.Contains(column2, "Yes") {
				// AbandonedLinks = append(AbandonedLinks, colly.) ?????? maybe the way to extract parents link
				IsOriginal = true
			}
		}
	})
	c.Visit(url)

	return IsOriginal
}

func CheckAbandonOne(url string) bool {
	IsAbandoned := false
	c := colly.NewCollector()

	c.OnHTML("div.row > div.col > table tr", func(e *colly.HTMLElement) {
		column1 := e.DOM.Find("td:nth-child(1)").Text()
		if column1 == "Category" {
			column2 := e.DOM.Find("td:nth-child(2)").Text()
			IsAbandoned = strings.Contains(column2, "Abandonware")
			if IsAbandoned {
				IsAbandoned = strings.Contains(column2, "Operating Systems")
			}
		}
	})
	c.Visit(url)
	return IsAbandoned
}

func CheckRelease(Links []string) []string {
	var OriginalReleaseLinks []string
	IsOriginal := false
	c := colly.NewCollector()

	c.OnHTML("div.col-md-6 > table tr", func(e *colly.HTMLElement) {
		column1 := e.DOM.Find("td:nth-child(1)").Text()
		column2 := e.DOM.Find("td:nth-child(2)").Text()
		if strings.Contains(column1, "Original Release?") {
			if strings.Contains(column2, "Yes") {
				// AbandonedLinks = append(AbandonedLinks, colly.) ?????? maybe the way to extract parents link
				IsOriginal = true
			}
		}
	})
	for i, url := range Links {
		c.Visit(url)
		fmt.Printf("#%v. Checking link: %v\n", i, url)
		if IsOriginal {
			fmt.Printf("#%v. Original Release Link: %v\n", i, url)
			OriginalReleaseLinks = append(OriginalReleaseLinks, url)
		}
		IsOriginal = false
	}

	return OriginalReleaseLinks
}

func CheckAbandon(Links []string) []string {
	var AbandonWareLinks []string
	IsAbandoned := false
	c := colly.NewCollector()

	c.OnHTML("div.row > div.col > table tr", func(e *colly.HTMLElement) {
		column1 := e.DOM.Find("td:nth-child(1)").Text()
		if column1 == "Category" {
			column2 := e.DOM.Find("td:nth-child(2)").Text()
			IsAbandoned = strings.Contains(column2, "Abandonware")
			if IsAbandoned {
				IsAbandoned = strings.Contains(column2, "Operating Systems")
			}
		}
	})

	for i, url := range Links {
		c.Visit(url)
		fmt.Printf("#%v. Checking link: %v\n", i, url)
		if IsAbandoned {
			fmt.Printf("#%v. AbandonWare Link: %v\n", i, url)
			AbandonWareLinks = append(AbandonWareLinks, url)
		}
		IsAbandoned = false
	}

	return AbandonWareLinks
}
