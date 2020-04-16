package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"pricecrawl/adapters/trendyol"
)

func main() {
	trendyol.Crawl()
}
