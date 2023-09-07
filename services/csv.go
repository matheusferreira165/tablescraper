package services

import (
	"encoding/csv"
	"log"
	"os"
)

func GenerateCsv() {

	file, err := os.Create("./data/TableData.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	link := "https://www.w3schools.com/html/html_tables.asp"

	datatables, err := ExtractTable(link)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, datatable := range datatables.Headers {
		if err := writer.Write([]string{datatable}); err != nil {
			log.Fatal(err)
		}
	}

	for _, datatable := range datatables.Rows {
		if err := writer.Write([]string{datatable}); err != nil {
			log.Fatal(err)
		}
	}

}
