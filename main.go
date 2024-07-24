package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itsorganic/go-scraper/crawler"
)

func main() {
	app := fiber.New()
	var url string
	fmt.Println("Enter the name of the product: ")
	fmt.Scanln(&url)

	app.Get("/scrape", func(c *fiber.Ctx) error {
		products := crawler.ScrapeAmazon(url)
		return c.JSON(products)
	})
	log.Fatal(app.Listen(":3000"))
	fmt.Println("Listening on port :3000")
}
