package domain

type User struct {
	UID              string
	Name             string
	Email            string
	EmailVerified    bool
	PhotoUrl         string
	PhoneNumber      string
}

type UserTiny struct {
	UID         string
	Name        string
	PhotoUrl    string
}

type CoServiceUserTying struct {
	UID           string
	ServiceUID    string
}
