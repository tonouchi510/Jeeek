package repository

import "github.com/tonouchi510/Jeeek/domain"

type ActivityRepository interface {
	InsertAll(activities []*domain.Activity) (num int, err error)
	Insert(activity domain.Activity) error
}
