package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gocolly/colly/v2"
)

func main() {
	domain := os.Args[len(os.Args)-1]

	c := colly.NewCollector(colly.AllowedDomains(domain))

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		file := "out/" + r.FileName()
		os.MkdirAll(filepath.Dir(file), 0777)
		r.Save(file)
	})

	if len(os.Args) < 3 {
		c.Visit("https://" + domain)
	} else {
		for i := 1; i < len(os.Args)-1; i++ {
			c.Visit(os.Args[i])
		}
	}
}
