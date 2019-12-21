package model

const ExternalServiceCollection = "externalServices"

type ExternalServices struct {
	Services    []Service    `firestore:"services"`
}

type Service struct {
	ServiceName    string    `firestore:"name"`
	ServiceUID     string    `firestore:"uid"`
}
