package controller

import (
	"context"
	"firebase.google.com/go/auth"
	"google.golang.org/api/iterator"
	"log"

	user "github.com/tonouchi510/Jeeek/gen/user"
)

// User service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
	authClient	*auth.Client
}

// NewUser returns the User service implementation.
func NewUser(logger *log.Logger, authClient *auth.Client) user.Service {
	return &usersrvc{logger, authClient}
}

// 現在のセッションに紐付くユーザの情報を返します。
func (s *usersrvc) GetCurrentUser(ctx context.Context, p *user.SessionTokenPayload) (res *user.JeeekUser, view string, err error) {
	res = &user.JeeekUser{}
	view = "default"
	s.logger.Print("user.Get current user")

	verifiedToken, err := s.authClient.VerifyIDToken(ctx, *p.Token)
	u, err := s.authClient.GetUser(ctx, verifiedToken.UID)
	if err != nil {
		return
	}
	res = &user.JeeekUser{
		UserID: u.UID,
		UserName: u.DisplayName,
		EmailAddress: u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
	}
	return
}

// 現在のセッションに紐付くユーザー情報を更新します。
func (s *usersrvc) UpdateUser(ctx context.Context, p *user.UpdateUserPayload) (res *user.JeeekUser, view string, err error) {
	view = "default"
	s.logger.Print("user.Update user")

	verifiedToken, err := s.authClient.VerifyIDToken(ctx, *p.Token)

	params := (&auth.UserToUpdate{}).
		Email(*p.EmailAddress).
		PhoneNumber(*p.PhoneNumber).
		DisplayName(*p.UserName).
		PhotoURL(*p.PhotoURL)
	u, err := s.authClient.UpdateUser(ctx, verifiedToken.UID, params)
	if err != nil {
		return
	}
	res = &user.JeeekUser{
		UserID: u.UID,
		UserName: u.DisplayName,
		EmailAddress: u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
	}

	s.logger.Printf("Successfully updated user: %v\n", u)
	return
}

// ユーザーの一覧を返します。
func (s *usersrvc) ListUser(ctx context.Context, p *user.SessionTokenPayload) (res user.JeeekUserCollection, view string, err error) {
	res = user.JeeekUserCollection{}
	view = "default"
	s.logger.Print("user.List user")

	iter := s.authClient.Users(ctx, "")
	for {
		u, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			return nil, view, err
		}
		r := &user.JeeekUser{
			UserID: u.UID,
			UserName: u.DisplayName,
			EmailAddress: u.Email,
			PhoneNumber: u.PhoneNumber,
			PhotoURL: u.PhotoURL,
			EmailVerified: u.EmailVerified,
		}
		res = append(res, r)
	}
	return
}

// 指定したIDのユーザーの情報を返します。
func (s *usersrvc) GetUser(ctx context.Context, p *user.GetUserPayload) (res *user.JeeekUser, view string, err error) {
	view = "default"
	s.logger.Print("user.Get user")
	u, err := s.authClient.GetUser(ctx, p.UserID)
	if err != nil {
		return
	}
	res = &user.JeeekUser{
		UserID: u.UID,
		UserName: u.DisplayName,
		EmailAddress: u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
	}
	return
}

// 現在のセッションに紐づくユーザーを削除します。
func (s *usersrvc) DeleteUser(ctx context.Context, p *user.SessionTokenPayload) (err error) {
	s.logger.Print("user.Delete user")

	verifiedToken, err := s.authClient.VerifyIDToken(ctx, *p.Token)

	params := (&auth.UserToUpdate{}).Disabled(true)
	_, err = s.authClient.UpdateUser(ctx, verifiedToken.UID, params)
	if err != nil {
		return
	}
	s.logger.Printf("Successfully deleted user: %s\n", verifiedToken.UID)
	return
}
