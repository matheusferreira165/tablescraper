package services

import (
	"encoding/csv"
	"fmt"
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

	for _, datatable := range datatables.Table {

		var headersformated string
		var rowformated string

		for _, header := range datatable.Headers {
			headersformated += fmt.Sprintf("%s,", string(header))
		}

		for _, row := range datatable.Rows {
			rowformated += fmt.Sprintf("%v,", string(row))
		}

		if err := writer.Write([]string{"<!-- TABELA INICIO -->"}); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write([]string{headersformated}); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write([]string{rowformated}); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write([]string{"<!-- TABELA FIM -->"}); err != nil {
			log.Fatal(err)
		}

	}

}
