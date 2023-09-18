package models

import "github.com/uptrace/bun"

type Attraction struct {
	bun.BaseModel `         bun:"table:attractions"`
	Description   *string  `bun:"description"         json:"description"`
	MobilePhone   *string  `bun:"mobile_phone"        json:"mobilePhone"`
	LandLine      *string  `bun:"land_line"           json:"landLine"`
	Price         *float64 `bun:"price"               json:"price"`
	Name          string   `bun:"name,notnull"        json:"name"`
	ProfileID     int64    `bun:"profile_id"          json:"profileId"`
	CategoryID    int64    `bun:"category_id"         json:"categoryId"`
	ID            int64    `bun:"id,pk,autoincrement" json:"id"`
	Lat           float64  `bun:"lat,notnull"         json:"lat"`
	Lng           float64  `bun:"lng,notnull"         json:"lng"`
}
