package services

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/matheusferreira165/tablescraper/models"
)

func ExtractTable(link string) (models.TableData, error) {

	var tableData models.TableData

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

		s.Find("th").Each(func(_ int, s *goquery.Selection) {
			tableData.Headers = append(tableData.Headers, s.Text())
		})

		s.Find("td").Each(func(index int, s *goquery.Selection) {
			tableData.Rows = append(tableData.Rows, s.Text())
		})

	})

	return tableData, err
}
