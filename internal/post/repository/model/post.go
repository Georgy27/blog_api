package model

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int64        `db:"id"`
	Info      PostInfo     `db:""`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type PostInfo struct {
	Title            string `db:"title"`
	Content          string `db:"content"`
	ShortDescription string `db:"short_description"`
	BlogID           int64  `db:"blog_id"`
	BlogName         string `db:"blog_name"`
}

type UpdatePostInfo struct {
	BlogID           int64          `db:"blog_id"`
	Title            sql.NullString `db:"title"`
	ShortDescription sql.NullString `db:"short_description"`
	Content          sql.NullString `db:"content"`
}
