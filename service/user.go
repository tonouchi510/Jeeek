package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service/model"
)

type userService struct{
	ctx         context.Context
	authClient	*auth.Client
	fsClient	*firestore.Client
}

func NewUserService(ctx context.Context, authClient *auth.Client, fsClient *firestore.Client) repository.UserRepository {
	return &userService{ctx, authClient, fsClient}
}

func (s userService) GetUserByToken(token string) (res *domain.User, err error) {
	// tokenの検証 & パース
	verifiedToken, err := s.authClient.VerifyIDToken(s.ctx, token)
	if err != nil {
		return nil, err
	}

	u, err := s.authClient.GetUser(s.ctx, verifiedToken.UID)
	if err != nil {
		return nil, err
	}

	if u.Disabled {
		return nil, fmt.Errorf("error: user of uid=%s is disabled", u.UID)
	}

	res = &domain.User{
		UID:           u.UID,
		Name:          u.DisplayName,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		PhotoUrl:      u.PhotoURL,
		PhoneNumber:   u.PhoneNumber,
	}

	return res, nil
}

func (s userService) GetUserTinyByToken(token string) (res *domain.UserTiny, err error) {
	// tokenの検証 & パース
	verifiedToken, err := s.authClient.VerifyIDToken(s.ctx, token)
	if err != nil {
		return nil, err
	}

	u, err := s.authClient.GetUser(s.ctx, verifiedToken.UID)
	if err != nil {
		return nil, err
	}

	if u.Disabled {
		return nil, fmt.Errorf("error: user of uid=%s is disabled", u.UID)
	}

	res = &domain.UserTiny{
		UID:      u.UID,
		Name:     u.DisplayName,
		PhotoUrl: u.PhotoURL,
	}
	return res, nil
}

func (s userService) GetUserTinyByUID(uid string) (res *domain.UserTiny, err error) {
	u, err := s.authClient.GetUser(s.ctx, uid)
	if err != nil {
		return nil, err
	}

	if u.Disabled {
		return nil, fmt.Errorf("error: user of uid=%s is disabled", u.UID)
	}

	res = &domain.UserTiny{
		UID:      u.UID,
		Name:     u.DisplayName,
		PhotoUrl: u.PhotoURL,
	}
	return res, nil
}

func (s userService) GetFollowersByUID(uid string) (res []domain.UserTiny, err error) {
	dsnap, err := s.fsClient.Collection(model.UserCollection).Doc(uid).Get(s.ctx)
	if err != nil {
		return nil, err
	}
	var m model.Follows
	err = dsnap.DataTo(&m)
	if err != nil {
		return nil, err
	}

	for _, f := range m.Followers {
		a := &domain.UserTiny{
			UID: f.UID,
			Name: f.Name,
			PhotoUrl: f.PhotoUrl,
		}
		res = append(res, *a)
	}

	return res, nil
}
