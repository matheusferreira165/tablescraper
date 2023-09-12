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

		for _, header := range datatable.Headers {
			headersformated += fmt.Sprintf("%s,", string(header))
		}

		if err := writer.Write([]string{"<!-- TABELA INICIO -->"}); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write([]string{headersformated}); err != nil {
			log.Fatal(err)
		}

		for _, row := range datatable.Rows {
			if err := writer.Write(row); err != nil {
				log.Fatal(err)
			}
		}

		if err := writer.Write([]string{"<!-- TABELA FIM -->"}); err != nil {
			log.Fatal(err)
		}

	}

}
