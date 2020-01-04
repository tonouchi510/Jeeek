package model

const UserCollection = "users"

type Follows struct {
	Followers    []UserTiny  `firestore:"followers"`
	Followings   []UserTiny  `firestore:"followings"`
}

type UserTiny struct {
	UID         string  `firestore:"uid"`
	Name        string  `firestore:"name"`
	PhotoUrl    string  `firestore:"photoUrl"`
}
