package model

import "time"

type Blog struct {
	ID        int64
	Name      string
	Info      *BlogInfo
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BlogInfo struct {
	Name        string
	Description string
	WebsiteUrl  string
}
