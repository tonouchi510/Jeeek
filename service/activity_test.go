package service_test

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/stretchr/testify/suite"
	"github.com/tonouchi510/Jeeek/domain"
	"github.com/tonouchi510/Jeeek/domain/repository"
	"github.com/tonouchi510/Jeeek/service"
	"testing"
)

type ServiceActivityTestSuite struct {
	suite.Suite
	context     context.Context
	fsClient    *firestore.Client
	svc         repository.ActivityRepository

	testUserID         string
	testUserName       string
	testUserPhotoUrl   string
	activity    domain.Activity
}

func TestServiceActivity(t *testing.T) {
	suite.Run(t, new(ServiceActivityTestSuite))
}

func (suite *ServiceActivityTestSuite) SetupTest() {
	suite.context = ctx
	suite.fsClient = fsClient
	suite.svc = service.NewActivityService(suite.context, suite.fsClient)

	suite.testUserID = "z5y5f5xmzs95k6b3"
	suite.activity = domain.Activity{
		ID:         "abcdefghijklmn0",
		Category:   0,
		Content:    domain.Content{
			Subject: "testです",
			Url:     "https://google.com",
			Comment: "テストですよ。。。",
		},
		Rank:       0,
		Tags:       []string{"General", "Python"},
		UserTiny:   domain.UserTiny{
			UID:      suite.testUserID,
			Name:     suite.testUserName,
			PhotoUrl: suite.testUserPhotoUrl,
		},
	}
}

func (suite *ServiceActivityTestSuite) TearDownTest() {
}

func (suite *ServiceActivityTestSuite) TestInsertAndListActivityM() {
	require := suite.Require()
	assert := suite.Assert()

	err := suite.svc.Insert(suite.testUserID, suite.activity)
	require.Nil(err)

	activities, err := suite.svc.List(suite.testUserID)
	require.Nil(err)
	assert.Equal(suite.activity.ID, activities[0].ID)
	assert.Equal(suite.activity.Category, activities[0].Category)
	assert.Equal(suite.activity.Content, activities[0].Content)
	assert.Equal(suite.activity.Rank, activities[0].Rank)
	assert.Equal(suite.activity.UserTiny, activities[0].UserTiny)
}

func (suite *ServiceActivityTestSuite) TestXDeleteActivityM() {
	require := suite.Require()

	err := suite.svc.Delete(suite.testUserID, suite.activity.ID)
	require.Nil(err)
}
