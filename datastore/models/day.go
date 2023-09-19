package models

import "github.com/uptrace/bun"

type Day struct {
	ID          string       `bun:"id,pk,autoincrement"                         json:"id"`
	Attractions []Attraction `bun:"m2m:attractions_to_days,join:Day=Attraction" json:"attractions"`
	PlanID      int64        `bun:"plan_id,notnull"                             json:"planId"`
}
type AttractionsToDays struct {
	bun.BaseModel `            bun:"table:attractions_to_days"`
	Attraction    *Attraction `bun:"rel:belongs-to,join:attraction_id=id"`
	Day           *Day        `bun:"rel:belongs-to,join:day_id=id"`
	DayID         string      `bun:"day_id,pk"`
	AttractionID  int64       `bun:"attraction_id,pk"`
}
