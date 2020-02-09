package domain

import "time"

type Activity struct {
	ID           string
	Category     int
	Content      Content
	Rank         int
	Tags         []string
	Favorites    []string
	Gifts        []string
	UserTiny     UserTiny
	UpdatedAt    time.Time
}

type Content struct {
	Subject    string
	Url        string
	Comment    string
}
