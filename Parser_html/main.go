package main

import (
	"bufio"
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
)

func main() {
	//pageURL := "https://www.betaarchive.com/database/sitemap.php"
	//ScrapIntoFile(pageURL)

	AllLinks := ReadFromFile()
	SliceToFile(AllLinks)

	//ChosenLinks := ScrapPageURLs(pageURL)
	//SliceToFile(ChosenLinks)

	// AllLinks := ScrapPageURLs(pageURL)
	//AbandonWareLinks := CheckAbandon(AllLinks)
	//OriginalReleaseLinks := CheckRelease(AbandonWareLinks)
	//fmt.Println(OriginalReleaseLinks)
	//fmt.Println("------------Finished Checking Links------------")
	//SliceToFile(OriginalReleaseLinks)
	//
	//fmt.Println("Basic array length:", len(AllLinks))
	//fmt.Printf("Abandonware & Operating systems link array length: %v\n", len(AbandonWareLinks))
	//fmt.Printf("Final link array length: %v\n", len(OriginalReleaseLinks))
}

func ReadFromFile() []string {
	file, err := os.Open("AllLinks.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var List []string
	max := 1000
	for scanner.Scan() && max > 0 {
		max--
		url := scanner.Text()
		fmt.Printf("%v. Checking url : %v\n", max, url)
		if CheckAbandonOne(url) {
			if CheckOneRelease(url) {
				List = append(List, url)
				fmt.Println("---------Chosen one :", url)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return List
}

func ScrapIntoFile(url string) {
	file, err := os.Create("AllLinks.txt")
	if err != nil {
		fmt.Println("Error creating file")
	}

	c := colly.NewCollector()
	// fmt.Println("......Collector created......")

	c.OnHTML("url > loc", func(e *colly.HTMLElement) {
		url := e.DOM.Text()
		file.Write([]byte(url + "\n"))
	})
	c.Visit(url)
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
