package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
)

func main() {
	var word string
	fmt.Println("Введите слово по которому надо найти обозначение")
	fmt.Scanf("%s\n", &word)

	c := colly.NewCollector()

	c.OnHTML("div.search-result.highlight", func(h *colly.HTMLElement) {
		h.ForEach("p.truncate", func(i int, e *colly.HTMLElement) {
			fmt.Println(e.Text)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		logrus.Info("Visiting ", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		logrus.Info("Got a response from ", r.Request.URL)
	})

	c.OnError(func(r *colly.Response, err error) {
		logrus.Error(err.Error())
	})

	c.Visit("https://ozhegov.slovaronline.com/search?s=" + word)
}
