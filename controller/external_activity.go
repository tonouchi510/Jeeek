package controller

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/tonouchi510/Jeeek/service"
	"log"

	externalactivity "github.com/tonouchi510/Jeeek/gen/external_activity"
)

// ExternalActivity service example implementation.
// The example methods log the requests and return zero values.
type externalActivitysrvc struct {
	logger 		*log.Logger
	authClient	*auth.Client
	fsClient	*firestore.Client
}

// NewExternalActivity returns the ExternalActivity service implementation.
func NewExternalActivity(logger *log.Logger, authClient *auth.Client, fsClient *firestore.Client) externalactivity.Service {
	return &externalActivitysrvc{logger, authClient, fsClient}
}

// 指定したユーザのQiitaの記事投稿を取得する
func (s *externalActivitysrvc) FetchQiitaArticle(ctx context.Context, p *externalactivity.SessionTokenPayload) (err error) {
	s.logger.Print("externalActivity.Fetch qiita article")

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
func (s *externalActivitysrvc) PickOutPastActivityOfQiita(ctx context.Context, p *externalactivity.SessionTokenPayload) (err error) {
	s.logger.Print("externalActivity.Pick out past activity of qiita")
	return
}
