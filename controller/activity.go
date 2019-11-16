package controller

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/tonouchi510/Jeeek/service"
	"log"
	"time"

	activity "github.com/tonouchi510/Jeeek/gen/activity"
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
func (s *activitysrvc) FetchQiitaArticleByQiitaUserID(ctx context.Context, p *activity.GetActivityPayload) (err error) {
	s.logger.Print("activity.Fetch qiita article by qiita-user-id")

	// ユーザ情報の取得
	userService := service.NewUserService(ctx, s.authClient)
	user, err := userService.GetUserTinyByToken(*p.Token)
	if err != nil {
		s.logger.Print(err)
		return
	}

	// Qiita記事の取得
	qiitaService := service.NewQiitaService()
	res, err := qiitaService.GetArticleByUserId(p.UserID)
	if err != nil {
		s.logger.Print(err)
		return
	}
	res.User = *user

	// Activityの追加
	activityService := service.NewActivityService(ctx, s.fsClient)
	err = activityService.Insert(*res)
	if err != nil {
		s.logger.Print(err)
	}

	return err
}

// qiita連携済みユーザのqiitaでのアクティビティ更新を行うジョブ
func (s *activitysrvc) BatchJobMethodToRefreshQiitaActivity(ctx context.Context) (err error) {
	s.logger.Print("activity.Batch job method to refresh qiita activity")

	coService := service.NewCoServiceService(ctx, s.fsClient)
	userService := service.NewUserService(ctx, s.authClient)
	qiitaService := service.NewQiitaService()
	activityService := service.NewActivityService(ctx, s.fsClient)

	userList, err := coService.GetQiita()
	for _, u := range userList {
		time.Sleep(1 * time.Second)

		// ユーザ情報の取得
		user, err := userService.GetUserTinyByUID(u.UID)
		if err != nil {
			s.logger.Print(err)
			continue
		}

		// Qiita記事の取得
		res, err := qiitaService.GetArticleByUserId(u.ServiceUID)
		if err != nil {
			s.logger.Print(err)
			continue
		}
		res.User = *user

		// Activityの追加
		err = activityService.Insert(*res)
		if err != nil {
			s.logger.Print(err)
			continue
		}
		s.logger.Printf("Success: refresh qiita activity for user uid = %s", u.UID)
	}

	return
}

// サービス連携以前のqiita記事投稿を反映させる
func (s *activitysrvc) PickOutPastActivityOfQiita(ctx context.Context, p *activity.GetActivityPayload) (err error) {
	s.logger.Print("activity.Pick out past activity of qiita")
	return
}
