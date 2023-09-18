package models

import "github.com/uptrace/bun"

type Profile struct {
	bun.BaseModel `              bun:"table:profiles"`
	FirstName     string        `bun:"first_name"                      json:"firstName"`
	LastName      string        `bun:"last_name"                       json:"lastName"`
	Email         string        `bun:"email"                           json:"email"`
	Interests     []*Category   `bun:"rel:has-many,join:id=profile_id" json:"interests"`
	Country       *Country      `bun:"rel:has-one,join:id=profile_id"  json:"country"`
	Favourites    []*Attraction `bun:"rel:has-many,join:id=profile_id" json:"attractions"`
	ID            int64         `bun:"id,pk,autoincrement"             json:"id"`
}
