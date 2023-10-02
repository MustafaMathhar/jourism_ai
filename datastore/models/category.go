package models

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `             bun:"table:categories"`
	Description   *string      `bun:"description"                                            json:"description,"`
	Name          string       `bun:"name"                                                   json:"name,"`
	Attractions   []Attraction `bun:"m2m:attractions_to_categories,join:Category=Attraction" json:"attractions"`
	ProfileID     int64        `bun:"profile_id"                                             json:"profileId"`
	ID            int64        `bun:"id,pk,autoincrement"                                    json:"id,"`
}

type AttractionsToCategories struct {
	bun.BaseModel `            bun:"table:attractions_to_categories"`
	Attraction    *Attraction `bun:"rel:belongs-to,join:attraction_id=id"`
	Category      *Category   `bun:"rel:belongs-to,join:category_id=id"`
	CategoryID    int64       `bun:"category_id,pk"`
	AttractionID  int64       `bun:"attraction_id,pk"`
}
