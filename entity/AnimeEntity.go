package entity

import (
	"database/sql"
)

type AnimeEntity struct {
	Id          int            `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	Episodes    sql.NullInt32  `json:"episodes"`
	Aired       sql.NullTime   `json:"aired"`
	Finished    sql.NullTime   `json:"finished"`
}
