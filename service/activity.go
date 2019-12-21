package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type activityService struct{
	ctx         context.Context
	fsClient	*firestore.Client
}

func NewActivityService(ctx context.Context, client *firestore.Client) repository.ActivityRepository {
	return &activityService{ctx, client}
}

func (s activityService) Insert(activity domain.Activity) (err error) {
	snapshot, err := s.fsClient.Collection(model.UserCollection).Doc(activity.User.UID).
		Collection(model.ActivityCollection).Doc(activity.ID).Get(s.ctx)
	if err != nil && grpc.Code(err) != codes.NotFound {
		return
	}
	if snapshot.Exists() {
		return fmt.Errorf("error: activity id=%s is already exist", activity.ID)
	}

	data := &model.Activity{
		Category:  activity.Category,
		Content:   model.Content{
			Subject: activity.Content.Subject,
			Url: activity.Content.Url,
			Comment: activity.Content.Comment,
		},
		Rank:      activity.Rank,
		Tags:      activity.Tags,
		User:      model.User{
			UID: activity.User.UID,
			Name: activity.User.Name,
			PhotoUrl: activity.User.PhotoUrl,
		},
		UpdatedAt: time.Now(),
	}
	_, err = s.fsClient.Collection(model.UserCollection).Doc(activity.User.UID).
		Collection(model.ActivityCollection).Doc(activity.ID).Set(s.ctx, data)

	// TODO:フォロワータイムラインへの反映ジョブのパブリッシュ

	return err
}
