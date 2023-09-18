package models

import "github.com/uptrace/bun"

type Profile struct {
	bun.BaseModel `              bun:"table:profiles"`
	FirstName     string        `bun:"first_name"                        json:"firstName"`
	LastName      string        `bun:"last_name"                         json:"lastName"`
	Email         string        `bun:"email"                             json:"email"`
	Interests     []*Category   `bun:"rel:has-many,join:id=profile_id"   json:"interests"`
	Country       *Country      `bun:"rel:belongs-to,join:country_id=id" json:"country"`
	CountryID     *int32        `bun:"country_id" json:"-"`
	Favourites    []*Attraction `bun:"rel:has-many,join:id=profile_id"   json:"favourites"`
	ID            int64         `bun:"id,pk,autoincrement"               json:"id"`
}
