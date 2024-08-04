package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/itsorganic/go-scraper/crawler"
)

func main() {
	app := fiber.New()
	product := flag.String("product", "", "search a product by giving keyword")
	flag.Parse()
	if *product == "" {
		log.Fatal("Flag is required")
	} else {
		app.Get("/", func(c *fiber.Ctx) error {
			products := crawler.ScrapeAmazon(*product)
			return c.JSON(products)
		})
		log.Fatal(app.Listen(":3000"))
		fmt.Println("Listening on port :3000")
	}
}
