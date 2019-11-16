package repository

import "github.com/tonouchi510/Jeeek/domain"

type QiitaRepository interface {
	GetArticleByUserId(userID string) (res *domain.Activity, err error)
}
