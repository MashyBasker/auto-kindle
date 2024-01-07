package scraper

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

type Book struct {
	Name		string
	Link 		string
	Author		string
	Publisher	string
}

func ScrapeBookNames(keyword string) []Book {
	c := colly.NewCollector(
		colly.AllowedDomains("singlelogin.re"),
	)
	var books []Book
	c.OnHTML("div.resItemBoxBooks", func(h *colly.HTMLElement) {
		var book Book
		h.ForEach("h3", func (_ int, h3 *colly.HTMLElement)  {
			bookName := strings.TrimSpace(h3.Text)
			link := h.ChildAttr("a", "href")
			book.Name = bookName
			book.Link = "https://singlelogin.re/"+link
		})
		h.ForEach("div.authors", func(_ int, h *colly.HTMLElement) {
			book.Author = h.Text
		})
		h.ForEach("[title=Publisher]", func(_ int, h *colly.HTMLElement) {
			fmt.Println("Publisher name: ", h.Text)
			book.Publisher = h.Text
		})
		books = append(books, book)
	})
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL: ", r.Request.URL, " failed with response: ", r, "\nError: ", err)
	})
	err := c.Visit("https://singlelogin.re/s/"+keyword)
	if err != nil {
		log.Fatal(err)
	}
	return books
}