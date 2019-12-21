package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service/model"
)

type externalServiceService struct{
	ctx         context.Context
	fsClient	*firestore.Client
}

func NewExternalServiceService(ctx context.Context, client *firestore.Client) repository.ExternalServiceRepository {
	return &externalServiceService{ctx, client}
}

func (s externalServiceService) GetQiita(uid string) (res *domain.ExternalServiceUser, err error) {
	dsnap, err := s.fsClient.Collection(model.ExternalServiceCollection).Doc(uid).Get(s.ctx)
	if err != nil {
		return nil, err
	}
	var users model.ExternalServices
	err = dsnap.DataTo(&users)
	if err != nil {
		return nil, err
	}

	for _, s := range users.Services {
		if s.ServiceName == "qiita" {
			res = &domain.ExternalServiceUser{
				ServiceName: s.ServiceName,
				ServiceUID: s.ServiceUID,
			}
		}
	}

	return res, nil
}
