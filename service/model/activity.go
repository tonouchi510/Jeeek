package model

import (
	"time"
)

const ActivityCollection = "activities"

type Activity struct {
	ID           string      `firestore:"-"`
	Category     int         `firestore:"category"`
	Content      Content     `firestore:"content"`
	Rank         int         `firestore:"rank"`
	Tags         []string    `firestore:"tags"`
	User         User        `firestore:"user"`
	UpdatedAt    time.Time   `firestore:"updatedAt"`
}

type Content struct {
	Subject    string    `firestore:"subject"`
	Url        string	 `firestore:"url"`
	Comment    string    `firestore:"comment"`
}

type User struct {
	UID         string    `firestore:"uid"`
	Name        string    `firestore:"name"`
	PhotoUrl    string    `firestore:"photoUrl"`
}
