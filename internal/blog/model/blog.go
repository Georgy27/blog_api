package model

import (
	"database/sql"
	"time"
)

type Blog struct {
	ID        int64
	Info      BlogInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type BlogInfo struct {
	Name        string
	Description string
	WebsiteUrl  string
}

type UpdateBlogInfo struct {
	Name        sql.NullString
	Description sql.NullString
	WebsiteUrl  sql.NullString
}
