package models

import "time"

type Product struct {
	ID            int32     `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Brand         string    `json:"brand"`
	Title         string    `json:"title"`
	PriceNormal   string    `json:"price_normal"`
	PriceDiscount string    `json:"price_discount"`
	URL           string    `json:"url"`
	Site          Site      `json:"site" gorm:"foreignkey:SiteID"`
	SiteID        int32     `json:"-"`
	UpdatedAt     time.Time `json:"updated_at"`
	CreatedAt     time.Time `json:"created_at"`
}
