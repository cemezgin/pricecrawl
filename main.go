package main

import (
	"crawltest/adapters/trendyol"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

func main() {
	trendyol.Crawl()
}
