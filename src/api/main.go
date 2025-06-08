package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type obra struct {
	Titulo string
	Ano    string
	Link   string
}
type camarada struct {
	nome   string
	vida   string
	resumo string
	obras  []obra
}

func main() {
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only domains: hackerspaces.org, wiki.hackerspaces.org
		colly.AllowedDomains("www.marxists.org"),
	)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)

	})
	c.OnHTML("tr td", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		//	fmt.Println( e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.marxists.org/portugues/lenin/index.htm")
}
