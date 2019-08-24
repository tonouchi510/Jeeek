package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/tonouchi510/Jeeek/gen/admin"
	"io/ioutil"
	"log"
	"net/http"
)

// Admin service example implementation.
// The example methods log the requests and return zero values.
type adminsrvc struct {
	logger *log.Logger
	authClient 	*auth.Client
}

// NewAdmin returns the Admin service implementation.
func NewAdmin(logger *log.Logger, authClient *auth.Client) admin.Service {
	return &adminsrvc{logger, authClient}
}

// admin apiのhealth-check
func (s *adminsrvc) AdminHealthCheck(ctx context.Context, p *admin.SessionTokenPayload) (res *admin.JeeekHealthcheck, err error) {
	res = &admin.JeeekHealthcheck{}
	s.logger.Print("admin.admin health-check")
	res.Result = "OK"
	return
}

// admin権限のトークンを取得します．
func (s *adminsrvc) AdminSignin(ctx context.Context, p *admin.AdminSignInPayload) (res *admin.JeeekAdminSignin, err error) {
	res = &admin.JeeekAdminSignin{}
	s.logger.Print("admin.admin signin")

	claims := map[string]interface{}{"admin": true}
	token, err := s.authClient.CustomTokenWithClaims(ctx, p.UID, claims)
	if err != nil {
		log.Fatalf("error minting custom token: %v\n", err)
	}

	req, err := json.Marshal(map[string]interface{}{
		"token":             token,
		"returnSecureToken": true,
	})
	if err != nil {
		return
	}

	// firebase_apikeyは一応晒して言いそうなので直書きしてる
	path := fmt.Sprintf("https://www.googleapis.com/identitytoolkit/v3/relyingparty/verifyCustomToken?key=%s",
		"AIzaSyB62WBiA_JWszHIRl7FrGFwK947_TwL0xo")
	r, err := http.Post(path, "application/json", bytes.NewBuffer(req))
	if err != nil {
		return
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http status code: %d", r.StatusCode)
	}
	resp, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	var resBody struct {
		IDToken string `json:"idToken"`
	}
	if err = json.Unmarshal(resp, &resBody); err != nil {
		return
	}
	res.Token = resBody.IDToken
	return
}
