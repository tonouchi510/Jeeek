package repository

import "github.com/tonouchi510/Jeeek/domain"

type ExternalServiceRepository interface {
	ListServiceAccounts(uid string) (res []*domain.ExternalServiceUser, err error)
	GetQiitaAccount(uid string) (res *domain.ExternalServiceUser, err error)
}
