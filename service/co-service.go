package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service/model"
)

type coServiceService struct{
	ctx         context.Context
	fsClient	*firestore.Client
}

func NewCoServiceService(ctx context.Context, client *firestore.Client) repository.CoServiceRepository {
	return &coServiceService{ctx, client}
}

func (s coServiceService) GetQiita() ([]*domain.CoServiceUserTying, error) {
	dsnap, err := s.fsClient.Collection(model.CoServiceCollection).Doc("qiita").Get(s.ctx)
	if err != nil {
		return nil, err
	}
	var users model.CoServiceUsers
	err = dsnap.DataTo(&users)
	if err != nil {
		return nil, err
	}

	var res []*domain.CoServiceUserTying
	for _, u := range users.Users {
		r := domain.CoServiceUserTying{
			UID: u.UID,
			ServiceUID: u.ServiceUID,
		}
		res = append(res, &r)
	}

	return res, nil
}
