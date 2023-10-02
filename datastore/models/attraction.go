package models

import "github.com/uptrace/bun"

type Attraction struct {
	bun.BaseModel `         bun:"table:attractions"`
	Location      *string  `bun:"location"          json:"location"`
	Description   *string  `bun:"description"       json:"description"`
	MobilePhone   *string  `bun:"mobile_phone"      json:"mobilePhone"`
	LandLine      *string  `bun:"land_line"         json:"landLine"`
	Price         *float64 `bun:"price"             json:"price"`
	BannerURL     *string  `bun:"banner_url"        json:"bannerUrl"`
	Name          string   `bun:"name,notnull"      json:"name"`

	ID  int64   `bun:"id,pk,autoincrement" json:"id"`
	Lat float64 `bun:"lat,notnull"         json:"lat"`
	Lng float64 `bun:"lng,notnull"         json:"lng"`
}
