package repository

import "github.com/tonouchi510/Jeeek/domain"

type ExternalServiceRepository interface {
	GetQiita(uid string) (res *domain.ExternalServiceUser, err error)
}
