package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("r-b-a.de"))

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

	c.Visit("https://r-b-a.de/")
}
