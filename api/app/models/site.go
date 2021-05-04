package models

import "time"

type SalesSite struct {
	Name      string    `json:"name,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SkillMarket struct {
	Name      string    `json:"name,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Sns struct {
	Name      string    `json:"name,omitempty"`
	Url       string    `json:"url,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
