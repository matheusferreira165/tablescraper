package services

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/matheusferreira165/tablescraper/models"
)

func ExtractTable(link string) (models.Table, error) {

	var table models.Table

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
		var tableData models.TableData

		s.Find("th").Each(func(_ int, s *goquery.Selection) {
			tableData.Headers = append(tableData.Headers, s.Text())
		})

		s.Find("td").Each(func(index int, s *goquery.Selection) {
			tableData.Rows = append(tableData.Rows, s.Text())
		})

		table.Table = append(table.Table, tableData)
	})

	return table, err
}
