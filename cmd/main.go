package main

import "github.com/matheusferreira165/tablescraper/services"

func main() {
	link := "https://www.espn.com.br/nba/calendario"

	services.ExtractTable(link)

}
