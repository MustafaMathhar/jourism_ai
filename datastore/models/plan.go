package models

import (
	"time"

	"github.com/uptrace/bun"
)

type Plan struct {
	bun.BaseModel `bun:"table:plans"`
	ProfileID     *int64     `bun:"profile_id"                   json:"profileId"`
	StartDate     *time.Time `bun:"start_date"                   json:"startDate"`
	EndDate       *time.Time `bun:"end_date"                     json:"endDate"`
	Name          string     `bun:"name,notnull"                 json:"name"`
	Days          []*Day     `bun:"rel:has-many,join:id=plan_id" json:"days"`
	ID            int64      `bun:"id,pk,autoincrement"          json:"id"`
}
