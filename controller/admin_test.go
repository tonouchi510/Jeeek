package controller_test

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"firebase.google.com/go/auth"
	"github.com/stretchr/testify/suite"
	"github.com/tonouchi510/Jeeek/gen/http/admin/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AdminControllerTestSuite struct {
	suite.Suite
	fsClient   *firestore.Client
	authClient *auth.Client

	token      string
	adminToken string
}

func TestAdminController(t *testing.T) {
	suite.Run(t, new(AdminControllerTestSuite))
}

func (suite *AdminControllerTestSuite) SetupTest() {
	suite.fsClient = fsClient
	suite.authClient = authClient

	suite.token = testToken
	suite.adminToken = adminToken
}

func (suite *AdminControllerTestSuite) TearDownTest() {
}

func (suite *AdminControllerTestSuite) TestAdminHealthCheck() {
	t := suite.T()
	require := suite.Require()
	assert := suite.Assert()

	path := server.AdminHealthCheckAdminPath()
	tests := []testCase{
		{
			description: "管理者権限でアクセスした場合、admin APIのhealth-checkができる",
			req: func() *http.Request {
				req := httptest.NewRequest("GET", path, nil)
				req.Header.Add("Authorization", suite.adminToken)
				return req
			},
			status: 200,
			condition: func(rr *httptest.ResponseRecorder, t *testing.T) {

				var body server.AdminHealthCheckResponseBody
				var err error
				err = json.Unmarshal(rr.Body.Bytes(), &body)
				require.Nil(err)
				assert.Equal("OK", body.Result)
			},
		},
		{
			description: "無効なセッション（トークン）でリクエストしたら401を返す",
			req: func() *http.Request {
				req := httptest.NewRequest("GET", path, nil)
				req.Header.Add("Authorization", "abcdefghijklmn")
				return req
			},
			status: 401,
			condition: func(rr *httptest.ResponseRecorder, t *testing.T) {
			},
		},
		{
			description: "セッションが確立されていない場合、401を返す",
			req: func() *http.Request {
				req := httptest.NewRequest("GET", path, nil)
				return req
			},
			status: 401,
			condition: func(rr *httptest.ResponseRecorder, t *testing.T) {
			},
		},
		{
			description: "管理者権限がないトークンでアクセスした場合、401を返す",
			req: func() *http.Request {
				req := httptest.NewRequest("GET", path, nil)
				req.Header.Add("Authorization", suite.token)
				return req
			},
			status: 401,
			condition: func(rr *httptest.ResponseRecorder, t *testing.T) {
			},
		},
	}
	for idx, test := range tests {
		runTestCondition(idx, test, t)
	}
}
