package repository

import "github.com/tonouchi510/Jeeek/domain"

type ActivityRepository interface {
	Insert(activity domain.Activity) error
}
