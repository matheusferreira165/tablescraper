package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func ExtractTable(link string) {

	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("table").Each(func(i int, s *goquery.Selection) {

		headers := s.Find("th").Each(func(_ int, s *goquery.Selection) {

			if s != nil {
				fmt.Print(s.Text())
				fmt.Print(" ")
			}
		})

		fmt.Println()

		s.Find("td").Each(func(index int, s *goquery.Selection) {
			if s != nil {
				fmt.Print(s.Text())
				fmt.Print(" ")

			}
			// Printing columns nicely

			if (index+1)%headers.Size() == 0 {
				fmt.Println()
			}
		})

	})
}
