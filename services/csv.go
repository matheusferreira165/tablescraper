package services

import (
	"encoding/csv"
	"log"
	"os"
)

func GenerateCsv(link string) (*os.File, error) {

	token := TokenGenerator()
	file, err := os.Create("./data/TableData_" + token + ".csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	datatables, err := ExtractTable(link)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, datatable := range datatables.Table {

		var headersformated []string

		for _, header := range datatable.Headers {
			headersformated = append(headersformated, header)
		}

		if err := writer.Write([]string{"<!-- TABELA INICIO -->"}); err != nil {
			log.Fatal(err)
		}

		if err := writer.Write(headersformated); err != nil {
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

	return file, err
}
