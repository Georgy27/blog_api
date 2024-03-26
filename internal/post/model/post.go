package model

import (
	"database/sql"
	"time"
)

type Post struct {
	ID        int64
	Info      PostInfo
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type PostInfo struct {
	Title            string
	Content          string
	ShortDescription string
	BlogID           int64
	BlogName         string
}

type ListPost struct {
	PagesCount int64
	Page       int64
	PageSize   int64
	TotalCount int64
	Post       []*Post
}

type UpdatePostInfo struct {
	BlogID           int64
	Title            sql.NullString
	ShortDescription sql.NullString
	Content          sql.NullString
}
