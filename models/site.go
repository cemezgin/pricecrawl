package models

type Site struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Domain   string `json:"domain"`
	IsActive bool   `json:"is_active"`
}
