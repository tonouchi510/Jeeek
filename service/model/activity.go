package model

import (
	"time"
)

const ActivityCollection = "timeline"

type Activity struct {
	ID           string      `firestore:"-"`
	Category     int         `firestore:"category"`
	Content      Content     `firestore:"content"`
	Rank         int         `firestore:"rank"`
	Tags         []string    `firestore:"tags"`
	Favorites    []string    `firestore:"favorites"`
	Gifts        []string    `firestore:"gifts"`
	UserTiny     UserTiny    `firestore:"userTiny"`
	UpdatedAt    time.Time   `firestore:"updatedAt"`
}

type Content struct {
	Subject    string    `firestore:"subject"`
	Url        string	 `firestore:"url"`
	Comment    string    `firestore:"comment"`
}
