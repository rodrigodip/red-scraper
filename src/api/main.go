package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"regexp"
	"strings"
)

type obra struct {
	Ano     string
	Mes     *string
	Titulo  string
	Link    string
	Content []string
}
type camarada struct {
	nome   string
	vida   string
	resumo string
	obras  []obra
}

func (o obra) printer() {
	if o.Mes != nil {
		fmt.Printf("Ano: %s - %s \n", o.Ano, *o.Mes)
	} else {
		fmt.Println("Ano:", o.Ano)
	}
	fmt.Printf("Título: %s \n", o.Titulo)
	fmt.Printf("Link: %s \n", o.Link)
	fmt.Printf("\n")
}

func takeContent(o obra) {
	link := o.Link
	c := colly.NewCollector()

	c.OnHTML("main", func(e *colly.HTMLElement) {
		fmt.Println(e.ChildText("p"))
	})
	c.Visit(link)
}

func main() {

	obras := []obra{}
	// Instantiate default collector
	c := colly.NewCollector(
		// Visit only this domain
		colly.AllowedDomains("www.marxists.org"),
	)

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)

	})
	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		temp := obra{}
		numb := 0 // só pra numerar a lista de obras na tela
		e.ForEach("tr", func(i int, e *colly.HTMLElement) {
			// handles data da obra separando Ano e Mês
			dataObra := e.ChildText("td:nth-of-type(1)")
			titObra := e.ChildText("td:nth-of-type(2)")
			linkObra := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))

			ano, mes, temMes := strings.Cut(dataObra, " - ") // Separa Ano do Mês
			temp.Ano = ano
			temp.Titulo = titObra
			temp.Link = linkObra

			// Verifica validade do conteúdo
			pattern := `^[0-9]{4}` // Apenas conteúdos que possuem a estrutura Ano - Titulo - Link são coletadas
			isValido := regexp.MustCompile(pattern)

			if isValido.MatchString(ano) { //valida conteudo
				numb++ // numera as obras apenas print
				if temMes {
					temp.Mes = &mes
				} else {
					temp.Mes = nil
				}
				temp.Titulo = titObra
				temp.Link = linkObra

				obras = append(obras, temp)

				//fmt.Printf("> %d -----------------------\n", numb)
				//obras[numb-1].printer()

			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL.String())
	})

	// Start scraping
	authorURL := "https://www.marxists.org/portugues/lenin/index.htm"
	c.Visit(authorURL)
	takeContent(obras[0])
}
