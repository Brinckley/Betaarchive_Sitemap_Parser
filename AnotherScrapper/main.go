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
	//CheckURLs(pageURL)
	ScrapURLs(pageURL, 0)
}

func CheckURLs(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("url > loc").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		// title := s.Find("a").Text()
		link := s.Text()
		fmt.Printf("Review %d: %s\n", i, link)
	})
}

func ScrapURLs(url string, index int) {
	file, err := os.OpenFile("Links.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("FULL RESTART")
		if index > 15 {
			log.Fatal(err)
		}
		index++
		ScrapURLs(url, index)
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
		if i >= 39616 {
			if ScrapAbandon(link, 0) {
				if ScrapRelease(link, 0) {
					file.WriteString(link + "\n")
					fmt.Println("---Selected :", link)
				}
			}
		}
	})
}

func ScrapAbandon(url string, index int) bool {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("...Restarting...")
		if index > 30 {
			log.Fatal(err)
		}
		index++
		return ScrapAbandon(url, index)
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

func ScrapRelease(url string, index int) bool {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println("...Restarting...")
		if index > 30 {
			log.Fatal(err)
		}
		index++
		return ScrapRelease(url, index)
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
