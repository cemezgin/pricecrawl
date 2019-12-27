package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
	"pricecrawl.tech/adapters/trendyol"
)

func main() {
	trendyol.Crawl()
}
