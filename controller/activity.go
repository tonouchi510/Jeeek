package controller

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/tonouchi510/Jeeek/gen/activity"
	"github.com/tonouchi510/Jeeek/service"
	"log"
)

// Activity service example implementation.
// The example methods log the requests and return zero values.
type activitysrvc struct {
	logger 		*log.Logger
	authClient	*auth.Client
	fsClient	*firestore.Client
}

// NewActivity returns the Activity service implementation.
func NewActivity(logger *log.Logger, authClient *auth.Client, fsClient *firestore.Client) activity.Service {
	return &activitysrvc{logger, authClient, fsClient}
}

// 指定したユーザのQiitaの記事投稿を取得する
func (s *activitysrvc) FetchQiitaArticle(ctx context.Context, p *activity.SessionTokenPayload) (err error) {
	s.logger.Print("activity.Fetch qiita article")

	// ユーザ情報の取得
	userService := service.NewUserService(ctx, s.authClient)
	user, err := userService.GetUserTinyByToken(*p.Token)
	if err != nil {
		s.logger.Print(err)
		return
	}

	// QiitaアカウントIDの取得
	externalService := service.NewExternalServiceService(ctx, s.fsClient)
	res, err := externalService.GetQiita(user.UID)
	if err != nil {
		s.logger.Print(err)
		return
	}
	if res == nil {
		s.logger.Print(err)
		return
	}
	serviceUid := res.ServiceUID

	// Qiita記事の取得
	qiitaService := service.NewQiitaService()
	activities, err := qiitaService.GetArticleByUserId(serviceUid)
	if err != nil {
		s.logger.Print(err)
		return
	}
	activities.User = *user

	// Activityの追加
	activityService := service.NewActivityService(ctx, s.fsClient)
	err = activityService.Insert(*activities)
	if err != nil {
		s.logger.Print(err)
	}

	return err
}

// サービス連携以前のqiita記事投稿を反映させる
func (s *activitysrvc) PickOutPastActivityOfQiita(ctx context.Context, p *activity.SessionTokenPayload) (err error) {
	s.logger.Print("activity.Pick out past activity of qiita")
	return
}
