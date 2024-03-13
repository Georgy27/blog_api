package model

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID        int64        `db:"id"`
	Info      BlogInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type BlogInfo struct {
	Name        string `db:"name"`
	Description string `db:"description"`
	WebsiteUrl  string `db:"website_url"`
}

type UpdateBlogInfo struct {
	Name        sql.NullString `db:"name"`
	Description sql.NullString `db:"description"`
	WebsiteUrl  sql.NullString `db:"website_url"`
}
