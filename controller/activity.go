package controller

import (
	"cloud.google.com/go/firestore"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
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

// タイムラインへの書き込みを行う
func (s *activitysrvc) ReflectionActivity(ctx context.Context, p *activity.ActivityPostPayload) (err error) {
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
