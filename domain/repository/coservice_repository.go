package repository

import "github.com/tonouchi510/Jeeek/domain"

type CoServiceRepository interface {
	GetQiita() ([]*domain.CoServiceUserTying, error)
}
