package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeBookNames(keyword string) {
	c := colly.NewCollector(
		colly.AllowedDomains("singlelogin.re"),
	)
	
	c.OnHTML("div.resItemBoxBooks", func(h *colly.HTMLElement) {
		h.ForEach("h3", func (_ int, h3 *colly.HTMLElement)  {
			text := strings.TrimSpace(h3.Text)
			fmt.Println("Book name: ", text)
			link := h.ChildAttr("a", "href")
			fmt.Println("Book link: https://singlelogin.re"+link)
		})
		h.ForEach("div.authors", func(_ int, h *colly.HTMLElement) {
			fmt.Println("Author name: ", h.Text)
		})
		h.ForEach("[title=Publisher]", func(_ int, h *colly.HTMLElement) {
			fmt.Println("Publisher name: ", h.Text)
		})
		fmt.Println("----------------------------------")
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL: ", r.Request.URL, " failed with response: ", r, "\nError: ", err)
	})
	err := c.Visit("https://singlelogin.re/s/"+keyword)
	if err != nil {
		log.Fatal(err)
	}

}