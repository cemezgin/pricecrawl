package trendyol

import (
	"crawltest/db"
	"crawltest/models"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Article struct {
	Title       string
	Name        string
	MarketPrice string
	SalePrice   string
	URL         string
}

func Crawl() {

	c := colly.NewCollector(
		colly.AllowedDomains("www.trendyol.com"),
	)

	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  ".*trendyol.*",
		Parallelism: 1,
		Delay:       1 * time.Second,
	})

	detailCollector := c.Clone()

	c.OnRequest(func(r *colly.Request) {

		fmt.Println("Visiting: ", r.URL.String())
	})

	c.OnHTML(`a[href]`, func(e *colly.HTMLElement) {
		foundURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Contains(foundURL, "butikdetay") {
			detailCollector.Visit(foundURL)
		} else {
			c.Visit(foundURL)
		}
	})

	detailCollector.OnHTML(`div.product-info-container`, func(e *colly.HTMLElement) {
		fmt.Println("Scraping Content ", e.Request.URL.String())
		article := Article{}
		article.URL = e.Request.URL.String()
		e.ForEach("div.info-box", func(_ int, box *colly.HTMLElement) {
			article.Title = box.ChildText("div.brand")
			article.Name = box.ChildText("div.name")
		})
		e.ForEach("div.price-container", func(_ int, box *colly.HTMLElement) {
			article.MarketPrice = box.ChildText("div.market-price")
			article.SalePrice = box.ChildText("div.sale-price")
		})

		product := models.Product{
			Brand:         article.Title,
			Title:         article.Name,
			PriceNormal:   article.MarketPrice,
			PriceDiscount: article.SalePrice,
			SiteID:        2,
			URL:           article.URL,
			CreatedAt:     time.Now(),
		}

		DB := db.Connect()
		create := DB.Create(&product)
		defer DB.Close()

		if create != nil {
			fmt.Println("Saved: ", article)
		}
	})

	c.Visit("https://www.trendyol.com")
}
