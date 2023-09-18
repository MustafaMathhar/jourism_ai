package models

import "github.com/uptrace/bun"

type Category struct {
	bun.BaseModel `              bun:"table:categories"`
	Description   *string       `bun:"description"                      json:"description,"`
	Name          string        `bun:"name"                             json:"name,"`
	Attractions   []*Attraction `bun:"rel:has-many,join:id=category_id" json:"attractions"`
	ProfileID     int64         `bun:"profile_id"                       json:"profileId"`
	ID            int64         `bun:"id,pk,autoincrement"              json:"id,"`
}
