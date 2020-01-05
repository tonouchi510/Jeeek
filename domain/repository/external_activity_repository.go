package repository

import "github.com/tonouchi510/Jeeek/domain"

type ExternalActivityRepository interface {
	ListActivityByServiceUID(uid string) (res []*domain.Activity, err error)
	// 多様な外部サービスとの連携に対応するために不便だがこの様なメソッドにしている
	GetRecentActivityByServiceUID(uid string, num int) (res []*domain.Activity, err error)
}
