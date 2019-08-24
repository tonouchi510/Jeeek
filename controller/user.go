package controller

import (
	"context"
	"log"

	user "github.com/tonouchi510/Jeeek/gen/user"
)

// User service example implementation.
// The example methods log the requests and return zero values.
type usersrvc struct {
	logger *log.Logger
}

// NewUser returns the User service implementation.
func NewUser(logger *log.Logger) user.Service {
	return &usersrvc{logger}
}

// 現在のセッションに紐付くユーザの情報を返します。
func (s *usersrvc) GetCurrentUser(ctx context.Context, p *user.SessionTokenPayload) (res *user.JeeekUser, view string, err error) {
	res = &user.JeeekUser{}
	view = "default"
	s.logger.Print("user.Get current user")
	return
}

// 現在のセッションに紐付くユーザー情報を更新します。
func (s *usersrvc) UpdateUser(ctx context.Context, p *user.UpdateUserPayload) (res *user.JeeekUser, view string, err error) {
	res = &user.JeeekUser{}
	view = "default"
	s.logger.Print("user.Update user")
	return
}

// ユーザーの一覧を返します。
func (s *usersrvc) ListUser(ctx context.Context, p *user.SessionTokenPayload) (res user.JeeekUserCollection, view string, err error) {
	view = "default"
	s.logger.Print("user.List user")
	return
}

// 指定したIDのユーザーの情報を返します。
func (s *usersrvc) GetUser(ctx context.Context, p *user.GetUserPayload) (res *user.JeeekUser, view string, err error) {
	res = &user.JeeekUser{}
	view = "default"
	s.logger.Print("user.Get user")
	return
}

// 現在のセッションに紐づくユーザーを削除します。
func (s *usersrvc) DeleteUser(ctx context.Context, p *user.SessionTokenPayload) (err error) {
	s.logger.Print("user.Delete user")
	return
}
