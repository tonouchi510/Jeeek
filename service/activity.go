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

type activityService struct {
	ctx         context.Context
	fsClient	*firestore.Client
}

func NewActivityService(ctx context.Context, client *firestore.Client) repository.ActivityRepository {
	return &activityService{ctx, client}
}

func (s activityService) InsertAll(uid string, activities []*domain.Activity) (success int, err error) {
	success = 0
	for _, activity := range activities {
		snapshot, err := s.fsClient.Collection(model.UserCollection).Doc(uid).
			Collection(model.ActivityCollection).Doc(activity.ID).Get(s.ctx)
		if err != nil && grpc.Code(err) != codes.NotFound {
			return
		}
		if snapshot.Exists() {
			// すでに保存済みの記事まで遡ったら抜ける
			continue
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
			Favorites: activity.Favorites,
			Gifts:     activity.Gifts,
			UserTiny:      model.UserTiny{
				UID: activity.UserTiny.UID,
				Name: activity.UserTiny.Name,
				PhotoUrl: activity.UserTiny.PhotoUrl,
			},
			UpdatedAt: time.Now(),
		}
		_, err = s.fsClient.Collection(model.UserCollection).Doc(uid).
			Collection(model.ActivityCollection).Doc(activity.ID).Set(s.ctx, data)

		if err != nil {
			return success, err
		}
		success++
	}

	return success, nil
}

func (s activityService) Insert(uid string, activity domain.Activity) (err error) {
	snapshot, err := s.fsClient.Collection(model.UserCollection).Doc(uid).
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
		UserTiny:      model.UserTiny{
			UID: activity.UserTiny.UID,
			Name: activity.UserTiny.Name,
			PhotoUrl: activity.UserTiny.PhotoUrl,
		},
		UpdatedAt: time.Now(),
	}
	_, err = s.fsClient.Collection(model.UserCollection).Doc(uid).
		Collection(model.ActivityCollection).Doc(activity.ID).Set(s.ctx, data)

	return
}
