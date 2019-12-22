package controller

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go/auth"
	"github.com/tonouchi510/Jeeek/domain/repository"
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

// セッションに紐づくユーザの連携済みサービスのアクティビティを取得する
func (s *externalActivitysrvc) RefreshActivitiesOfExternalServices(ctx context.Context, p *externalactivity.SessionTokenPayload) (err error) {
	s.logger.Print("externalActivity.Refresh activities of external services")

	userService := service.NewUserService(ctx, s.authClient)
	externalService := service.NewExternalServiceService(ctx, s.fsClient)
	activityService := service.NewActivityService(ctx, s.fsClient)

	// ユーザ情報の取得
	user, err := userService.GetUserTinyByToken(*p.Token)
	if err != nil {
		s.logger.Print(err)
		return
	}

	// 連携サービスのアカウントIDの取得
	accounts, err := externalService.ListServiceAccounts(user.UID)
	if err != nil || accounts == nil {
		s.logger.Print(err)
		return
	}

	for _, a := range accounts {
		serviceUid := a.ServiceUID

		var extService repository.ExternalActivityRepository
		if a.ServiceName == "qiita" {
			extService = service.NewQiitaService()
		} else {
			s.logger.Print("")
			continue
		}

		i, N := 0, 3
		for {
			// サービス毎のアクティビティを取得
			activities, err := extService.GetRecentActivityByServiceUID(serviceUid, i+N)
			if err != nil {
				s.logger.Print(err)
				return err
			}

			// 被らないn件をみる
			if len(activities)-i <= 0 {
				break
			} else if len(activities)-i > N {
				activities = activities[i:i+N]
			} else {
				activities = activities[i:]
			}

			// アクティビティの保存
			for _, activity := range activities {
				activity.User = *user
			}
			successes, err := activityService.InsertAll(activities)
			if err != nil {
				s.logger.Print(err)
			}
			if successes != N {
				break
			}
			i += N
		}
	}

	return nil
}

// セッションに紐づくユーザのQiitaの記事投稿を取得する
func (s *externalActivitysrvc) RefreshQiitaActivity(ctx context.Context, p *externalactivity.SessionTokenPayload) (err error) {
	s.logger.Print("externalActivity.Refresh qiita activity")

	userService := service.NewUserService(ctx, s.authClient)
	externalService := service.NewExternalServiceService(ctx, s.fsClient)
	qiitaService := service.NewQiitaService()
	activityService := service.NewActivityService(ctx, s.fsClient)

	// ユーザ情報の取得
	user, err := userService.GetUserTinyByToken(*p.Token)
	if err != nil {
		s.logger.Print(err)
		return
	}

	// QiitaアカウントIDの取得
	account, err := externalService.GetQiitaAccount(user.UID)
	if err != nil {
		s.logger.Print(err)
		return
	}
	if account == nil {
		s.logger.Print(err)
		return
	}
	serviceUid := account.ServiceUID

	// Qiita記事の取得
	i, N := 0, 3
	for {
		// サービス毎のアクティビティを取得
		activities, err := qiitaService.GetRecentActivityByServiceUID(serviceUid, i+N)
		if err != nil {
			s.logger.Print(err)
			return err
		}
		s.logger.Print(len(activities))

		// 被らないn件をみる
		if len(activities)-i <= 0 {
			break
		} else if len(activities)-i > N {
			activities = activities[i:i+N]
		} else {
			activities = activities[i:]
		}

		// アクティビティの保存
		for _, activity := range activities {
			activity.User = *user
		}
		successes, err := activityService.InsertAll(activities)
		if err != nil {
			s.logger.Print(err)
		}
		if successes != N {
			break
		}
		i += N
	}

	return nil
}

// サービス連携以前のqiita記事投稿を反映させる
func (s *externalActivitysrvc) PickOutPastActivityOfQiita(ctx context.Context, p *externalactivity.SessionTokenPayload) (err error) {
	s.logger.Print("externalActivity.Pick out past activity of qiita")

	// ユーザ情報の取得
	userService := service.NewUserService(ctx, s.authClient)
	user, err := userService.GetUserTinyByToken(*p.Token)
	if err != nil {
		s.logger.Print(err)
		return
	}

	// QiitaアカウントIDの取得
	externalService := service.NewExternalServiceService(ctx, s.fsClient)
	res, err := externalService.GetQiitaAccount(user.UID)
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
	activities, err := qiitaService.ListActivityByServiceUID(serviceUid)
	if err != nil {
		s.logger.Print(err)
		return
	}

	// アクティビティの保存
	activityService := service.NewActivityService(ctx, s.fsClient)
	for _, activity := range activities {
		activity.User = *user
	}
	_, err = activityService.InsertAll(activities)
	if err != nil {
		s.logger.Print(err)
	}

	return err
}
