package model

const CoServiceCollection = "CooperationServices"

type CoServiceUsers struct {
	Users    []UserTying    `firestore:"users"`
}

type UserTying struct {
	UID           string
	ServiceUID    string
}
