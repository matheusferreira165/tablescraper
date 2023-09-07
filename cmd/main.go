package main

import "github.com/matheusferreira165/tablescraper/services"

func main() {

	link := "https://www.w3schools.com/html/html_tables.asp"

	services.ExtractTable(link)

}
