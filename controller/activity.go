package controller

import (
	"cloud.google.com/go/firestore"
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/service"
	"log"

	"github.com/tonouchi510/Jeeek/gen/activity"
)

// Activity service example implementation.
// The example methods log the requests and return zero values.
type activitysrvc struct {
	logger *log.Logger
	authClient	*auth.Client
	fsClient	*firestore.Client
}

// NewActivity returns the Activity service implementation.
func NewActivity(logger *log.Logger, authClient *auth.Client, fsClient *firestore.Client) activity.Service {
	return &activitysrvc{logger, authClient, fsClient}
}

// 手動投稿用のAPI
func (s *activitysrvc) ManualActivityPost(ctx context.Context, p *activity.ActivityPostPayload) (err error) {
	s.logger.Print("activity.Manual activity post")

	userService := service.NewUserService(ctx, s.authClient, s.fsClient)
	activityService := service.NewActivityService(ctx, s.fsClient)

	pubsubClient, err := pubsub.NewClient(ctx, ProjectID)
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	topic := pubsubClient.Topic(TopicID)

	verifiedToken, err := s.authClient.VerifyIDToken(ctx, *p.Token)
	if err != nil {
		return err
	}
	follows, err := userService.GetFollowsByUID(verifiedToken.UID)
	if err != nil {
		s.logger.Print(err)
		return
	}

	data := &domain.Activity{
		ID:        p.Activity.ID,
		Category:  p.Activity.Category,
		Content:   domain.Content{
			Subject:  p.Activity.Content.Subject,
			Url:      *p.Activity.Content.URL,
			Comment:  *p.Activity.Content.Comment,
		},
		Rank:      p.Activity.Rank,
		Tags:      p.Activity.Tags,
		Favorites: p.Activity.Favorites,
		Gifts:     p.Activity.Gifts,
		UserTiny:  domain.UserTiny{
			UID:       p.Activity.UserTiny.UID,
			Name:      p.Activity.UserTiny.Name,
			PhotoUrl:  *p.Activity.UserTiny.PhotoURL,
		},
	}

	// 自分のタイムラインに追加
	err = activityService.Insert(verifiedToken.UID, *data)
	if err != nil {
		return err
	}

	// フォロワーのライムラインに反映
	bytes, _ := json.Marshal(*data)
	for _, f := range follows.Followers {
		result := topic.Publish(ctx, &pubsub.Message{
			Attributes: map[string]string{
				"uid": f.UID,
			},
			Data: bytes,
		})
		// Block until the result is returned and a server-generated
		// ID is returned for the published message.
		id, err := result.Get(ctx)
		if err != nil {
			s.logger.Print(err)
		}
		s.logger.Print("Published a message; msg ID: " + id)
	}

	return
}

// タイムラインへの書き込みを行う
func (s *activitysrvc) ReflectionActivity(ctx context.Context, p *activity.ActivityWriterPayload) (err error) {
	s.logger.Print("activity.Reflection activity")

	activityService := service.NewActivityService(ctx, s.fsClient)

	uid := p.Attributes[0].UID
	data := &domain.Activity{}
	err = json.Unmarshal(p.Data, data)
	if err != nil {
		return
	}
	err = activityService.Insert(*uid, *data)

	return
}
