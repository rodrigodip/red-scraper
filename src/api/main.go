package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
)

type obra struct {
	Titulo string
	Ano    string
	Mes    string
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
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		temp := obra{}

		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			// handles data da obra separando Ano e Mês
			dataObra := e.ChildText("td:nth-of-type(1)")
			ano, mes, temMes := strings.Cut(dataObra, " - ")
			temp.Ano = ano

			// Verifica validade do conteúdo
			pattern := `^[0-9]{4}` // Apenas conteúdos que possuem a estrutura Ano - Titulo - Link são coletadas
			isValido := regexp.MustCompile(pattern)

			if isValido.MatchString(ano) {
				if temMes {
					temp.Mes = mes
				} else {
					temp.Mes = "Indefinido"
				}
				fmt.Println("Ano :", temp.Ano, "Mês :", temp.Mes)
			}
			//titulo := e.ChildText("td:nth-of-type(2)")
			//link := e.ChildAttr("td", "a")

		})

		fmt.Println(temp.Titulo)
		fmt.Println(temp.Link)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.marxists.org/portugues/lenin/index.htm")
}
