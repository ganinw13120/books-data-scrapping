package repository

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly"

	"books-data-scrapping/model"
)

type bookRepository struct{}

type IBookRepository interface {
	GetBooksData(name string, books *[]model.Book) error
}

func NewBookRepository() bookRepository {
	return bookRepository{}
}

// Get books data from web scarpping
func (b bookRepository) GetBooksData(name string, books *[]model.Book) error {
	c := colly.NewCollector()
	bookCount := 0
	c.OnHTML("div.productitem", func(e *colly.HTMLElement) {
		e.ForEach("div.item-details", func(i int, element *colly.HTMLElement) {
			if element.ChildAttr("p.txt-normal", "title") == "" {
				return
			}
			book := model.Book{
				Name: element.ChildAttr("p.txt-normal", "title"),
			}
			element.ForEach("a.inline-block", func(j int, desc *colly.HTMLElement) {
				link := desc.Attr("href")
				if strings.Contains(link, "writer") {
					book.Author = desc.Text
				}
			})
			*books = append(*books, book)
			bookCount++
		})
	})
	iterator := 0
	c.OnHTML("img", func(e *colly.HTMLElement) {
		url := e.Attr("src")
		if strings.Contains(url, "resource/product") && iterator <= bookCount && bookCount != 0 {
			(*books)[iterator].ImageUrl = url
			iterator++
		}
	})

	c.Visit(fmt.Sprintf("%s%s", os.Getenv("WEB_PATH"), url.QueryEscape(name)))
	return nil
}
