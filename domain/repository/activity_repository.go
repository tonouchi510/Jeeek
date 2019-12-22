package repository

import "github.com/tonouchi510/Jeeek/domain"

type ActivityRepository interface {
	InsertAll(activities []*domain.Activity) error
	Insert(activity domain.Activity) error
}
