package repository

import "github.com/tonouchi510/Jeeek/domain"

type UserRepository interface {
	GetUserByToken(token string) (res *domain.User, err error)
	GetUserTinyByToken(token string) (res *domain.UserTiny, err error)
}
