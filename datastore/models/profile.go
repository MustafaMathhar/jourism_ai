package models

import "github.com/uptrace/bun"

type Profile struct {
	bun.BaseModel `             bun:"table:profiles"`
	FirstName     string       `bun:"first_name"                                         json:"firstName"`
	LastName      string       `bun:"last_name"                                          json:"lastName"`
	Email         string       `bun:"email"                                              json:"email"`
	Interests     []*Category  `bun:"rel:has-many,join:id=profile_id"                    json:"interests"`
	Country       *Country     `bun:"rel:belongs-to,join:country_id=id"                  json:"country"`
	CountryID     *int32       `bun:"country_id"                                         json:"-"`
	Favourites    []Attraction `bun:"m2m:profiles_to_attractions,join:Profile=Attraction" json:"favourites"`
	ID            int64        `bun:"id,pk,autoincrement"                                json:"id"`
}
type ProfilesToAttractions struct {
	bun.BaseModel `            bun:"table:profiles_to_attractions"`
	Attraction    *Attraction `bun:"rel:belongs-to,join:attraction_id=id"`
	Profile       *Profile    `bun:"rel:belongs-to,join:profile_id=id"`
	ProfileID     int64       `bun:"profile_id,pk"`
	AttractionID  int64       `bun:"attraction_id,pk"`
}
