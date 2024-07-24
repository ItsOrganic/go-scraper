package crawler

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Product struct {
	Name  string `json:"name"`
	Stars string `json:"stars"`
	Price string `json:"price"`
}

func ScrapeAmazon(pro string) []Product {
	c := colly.NewCollector(
		colly.AllowedDomains("www.amazon.in"),
	)

	product_url := "https://www.amazon.in/s?k=" + pro
	var products []Product
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Link of the page:", r.URL)
	})

	c.OnHTML("div.s-result-list.s-search-results.sg-row", func(h *colly.HTMLElement) {
		h.ForEach("div.a-section.a-spacing-base", func(_ int, h *colly.HTMLElement) {
			var product Product
			product.Name = h.ChildText("span.a-size-base-plus.a-color-base.a-text-normal")
			product.Stars = h.ChildText("span.a-icon-alt")
			product.Price = h.ChildText("span.a-price-whole")

			products = append(products, product)
		})
	})

	c.Visit(product_url)

	return products
}
