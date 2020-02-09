package service

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service/model"
	"google.golang.org/api/iterator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"log"
	"sort"
	"time"
)

type activityService struct {
	ctx         context.Context
	fsClient	*firestore.Client
}

func NewActivityService(ctx context.Context, client *firestore.Client) repository.ActivityRepository {
	return &activityService{ctx, client}
}

func (s activityService) List(uid string) (res []*domain.Activity, err error) {
	iter := s.fsClient.Collection(model.UserCollection).Doc(uid).
		Collection(model.ActivityCollection).Documents(s.ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var m model.Activity
		err = doc.DataTo(&m)
		if err != nil {
			return nil, err
		}

		res = append(res, &domain.Activity{
			ID:        doc.Ref.ID,
			Category:  m.Category,
			Content:   domain.Content{
				Subject:  m.Content.Subject,
				Url:      m.Content.Url,
				Comment:  m.Content.Comment,
			},
			Rank:      m.Rank,
			Tags:      m.Tags,
			Favorites: m.Favorites,
			Gifts:     m.Gifts,
			UserTiny:  domain.UserTiny{
				UID:       m.UserTiny.UID,
				Name:      m.UserTiny.Name,
				PhotoUrl:  m.UserTiny.PhotoUrl,
			},
			UpdatedAt: m.UpdatedAt,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].UpdatedAt.UnixNano() > res[j].UpdatedAt.UnixNano()
	})
	return
}

func (s activityService) InsertAll(uid string, activities []*domain.Activity) (success int, err error) {
	success = 0
	for _, activity := range activities {
		snapshot, err := s.fsClient.Collection(model.UserCollection).Doc(uid).
			Collection(model.ActivityCollection).Doc(activity.ID).Get(s.ctx)
		if err != nil && grpc.Code(err) != codes.NotFound {
			return success, err
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
	// 認証トークンに紐づくUIDを引数に受けるため、UIDの検証、存在チェックは外側のロジック
	snapshot, err := s.fsClient.Collection(model.UserCollection).Doc(uid).
		Collection(model.ActivityCollection).Doc(activity.ID).Get(s.ctx)
	if err != nil && grpc.Code(err) != codes.NotFound {
		return
	}
	if snapshot.Exists() {
		log.Printf("error: activity id=%s is already exist", activity.ID)
		return nil
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
		Favorites: []string{},
		Gifts:     []string{},
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

func (s activityService) Delete(uid string, activityID string) (err error) {
	_, err = s.fsClient.Collection(model.UserCollection).Doc(uid).
		Collection(model.ActivityCollection).Doc(activityID).Delete(s.ctx)

	return
}
