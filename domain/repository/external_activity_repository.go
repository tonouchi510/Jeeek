package repository

import "github.com/tonouchi510/Jeeek/domain"

type ExternalActivityRepository interface {
	GetRecentActivityByServiceUID(uid string, num int) (res []*domain.Activity, err error)
	ListActivityByServiceUID(uid string) (res []*domain.Activity, err error)
}
