package models

import "github.com/uptrace/bun"

type Country struct {
	bun.BaseModel `bun:"table:countries"`
	Currency  *string `bun:"currency"            json:"currency"`
	Code      string  `bun:"code,notnull"        json:"code"`
	Name      string  `bun:"name,notnull"        json:"name"`
	ProfileID int64   `bun:"profile_id"          json:"profileId"`
	ID        int32   `bun:"id,pk,autoincrement" json:"id"`
}
