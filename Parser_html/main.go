package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"os"
)

func main() {
	pageURL := "https://www.betaarchive.com/database/sitemap.php"
	//	SliceToFile([]string{pageURL, pageURL})

	ChosenLinks := ScrapPageURLs(pageURL)
	SliceToFile(ChosenLinks)

	//AllLinks := ScrapPageURLs(pageURL)
	//AbandonWareLinks := CheckAbandon(AllLinks)
	//OriginalReleaseLinks := CheckRelease(AbandonWareLinks)
	//fmt.Println(OriginalReleaseLinks)
	//fmt.Println("------------Finished Checking Links------------")
	//SliceToFile(OriginalReleaseLinks)

	//fmt.Println("Basic array length:", len(AllLinks))
	//fmt.Printf("Abandonware & Operating systems link array length: %v\n", len(AbandonWareLinks))
	//fmt.Printf("Final link array length: %v\n", len(OriginalReleaseLinks))
}

func ScrapPageURLs(url string) []string {
	var Links []string
	index := 0
	indexChosen := 0
	c := colly.NewCollector()
	// fmt.Println("......Collector created......")

	c.OnHTML("url > loc", func(e *colly.HTMLElement) {
		url := e.DOM.Text()
		fmt.Printf("%v. Checking url : %v\n", index, url)
		index++
		IsOrgignal := CheckOneRelease(url)
		if IsOrgignal {
			IsAbandoned := CheckAbandonOne(url)
			if IsAbandoned {
				fmt.Printf("___%v. Selected link : %v\n", indexChosen, e.DOM.Text())
				Links = append(Links, e.DOM.Text())
				indexChosen++
			}
		}

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
