package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	pageURL := "https://www.betaarchive.com/database/sitemap.php"

	ScrapURLs(pageURL)
}

func ScrapURLs(url string) {
	file, err := os.Create("Links.txt")
	if err != nil {
		fmt.Println("Error creating file")
	}

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	//if res.StatusCode != 200 {
	//	log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	//}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items

	doc.Find("url > loc").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		// title := s.Find("a").Text()
		link := s.Text()
		fmt.Printf("Review %d: %s\n", i, link)
		if ScrapAbandon(link) {
			if ScrapRelease(link) {
				file.Write([]byte(link + "\n"))
				fmt.Println("---Selected :", link)
			}
		}
	})
}

func ScrapAbandon(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	IsAbandon := false
	doc.Find("div.row > div.col > table tr:nth-of-type(1) > td:last-child").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		// title := s.Find("a").Text()
		IsAbandon = strings.Contains(s.Text(), "Abandonware")
		if IsAbandon {
			IsAbandon = strings.Contains(s.Text(), "Operating Systems")
		}

		// fmt.Printf("Review %d: %s\n", i, title)
	})

	return IsAbandon
}

func ScrapRelease(url string) bool {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	IsOriginal := false
	doc.Find("div.col-md-6 > table tr:nth-of-type(4) > td:last-child").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		// title := s.Find("a").Text()
		// fmt.Printf("Review %d: %s\n", i, title)
		if strings.Contains(s.Text(), "Yes") {
			IsOriginal = true
		}
	})

	return IsOriginal
}
