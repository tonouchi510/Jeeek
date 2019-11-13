package service

import (
	"context"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
)

type userService struct{
	ctx         context.Context
	authClient	*auth.Client
}

func NewUserService(ctx context.Context, client *auth.Client) repository.UserRepository {
	return &userService{ctx, client}
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
