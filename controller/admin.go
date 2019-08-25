package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"firebase.google.com/go/auth"
	"fmt"
	"github.com/tonouchi510/Jeeek/gen/admin"
	"google.golang.org/api/iterator"
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

// 新しいユーザーを登録します。
func (s *adminsrvc) AdminCreateNewUser(ctx context.Context, p *admin.AdminCreateUserPayload) (res *admin.JeeekUser, view string, err error) {
	res = &admin.JeeekUser{}
	view = "admin"
	s.logger.Print("admin.admin create new user")

	params := (&auth.UserToCreate{}).
		Email(p.EmailAddress).
		DisplayName(p.UserName).
		PhoneNumber(p.PhoneNumber).
		PhotoURL(p.PhotoURL).
		EmailVerified(false).
		Disabled(false)

	u, err := s.authClient.CreateUser(ctx, params)
	if err != nil{
		log.Fatalf("error creating user :%v\n", err)
	}

	res = &admin.JeeekUser{
		UserID: u.UID,
		UserName: u.DisplayName,
		EmailAddress: u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
		Disabled: &u.Disabled,
	}
	log.Printf("Successfully created user: %v\n", u)
	return
}

// 指定したユーザー情報を更新します。
func (s *adminsrvc) AdminUpdateUser(ctx context.Context, p *admin.AdminUpdateUserPayload) (res *admin.JeeekUser, view string, err error) {
	res = &admin.JeeekUser{}
	view = "admin"
	s.logger.Print("admin.admin update user")

	params := &auth.UserToUpdate{}
	if p.UserName != nil {
		params.DisplayName(*p.UserName)
	}
	if p.EmailAddress != nil {
		params.Email(*p.EmailAddress)
	}
	if p.EmailVerified != nil {
		params.EmailVerified(*p.EmailVerified)
	}
	if p.PhoneNumber != nil {
		params.PhoneNumber(*p.PhoneNumber)
	}
	if p.PhotoURL != nil {
		params.PhotoURL(*p.PhotoURL)
	}
	if p.Disabled != nil {
		params.Disabled(*p.Disabled)
	}

	u, err := s.authClient.UpdateUser(ctx, p.UserID, params)
	if err != nil{
		log.Fatalf("error updating user: %v\n", err)
	}
	log.Printf("Successfully update user: %v\n", u)
	res = &admin.JeeekUser{
		UserID: u.UID,
		UserName: u.DisplayName,
		EmailAddress: u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
		Disabled: &u.Disabled,
	}
	return
}

// ユーザーの一覧を返します。
func (s *adminsrvc) AdminListUser(ctx context.Context, p *admin.SessionTokenPayload) (res admin.JeeekUserCollection, view string, err error) {
	res = admin.JeeekUserCollection{}
	view = "admin"
	s.logger.Print("admin.admin list user")

	iter := s.authClient.Users(ctx, "")
	for {
		u, err := iter.Next()
		if err != nil {
			if err == iterator.Done{
				break
			}
			return nil, view, err
		}
		r := &admin.JeeekUser{
			UserID: u.UID,
			UserName: u.DisplayName,
			EmailAddress: u.Email,
			PhoneNumber: u.PhoneNumber,
			PhotoURL: u.PhotoURL,
			EmailVerified: u.EmailVerified,
			Disabled: &u.Disabled,
		}
		res = append(res, r)
	}
	return
}

// 指定したIDのユーザーの情報を返します。
func (s *adminsrvc) AdminGetUser(ctx context.Context, p *admin.GetUserPayload) (res *admin.JeeekUser, view string, err error) {
	res = &admin.JeeekUser{}
	view = "admin"
	s.logger.Print("admin.admin get user")
	u, err := s.authClient.GetUser(ctx, p.UserID)
	if err != nil{
		return
	}
	res = &admin.JeeekUser{
		UserID: u.UID,
		UserName: u.DisplayName,
		EmailAddress: u.Email,
		PhoneNumber: u.PhoneNumber,
		PhotoURL: u.PhotoURL,
		EmailVerified: u.EmailVerified,
		Disabled: &u.Disabled,
	}
	return
}

// 指定したユーザーを削除します。
func (s *adminsrvc) AdminDeleteUser(ctx context.Context, p *admin.AdminDeleteUserPayload) (err error) {
	s.logger.Print("admin.admin delete user")
	err = s.authClient.DeleteUser(ctx, p.UserID)
	if err != nil{
		log.Fatalf("error deleting user: %v\n", err)
	} else {
		log.Printf("Successfully deleted user: %v\n:", p.UserID)
	}
	return
}

// ユーザ数の統計情報を返す
func (s *adminsrvc) AdminUserStats(ctx context.Context, p *admin.SessionTokenPayload) (res *admin.JeeekUserStats, err error) {
	res = &admin.JeeekUserStats{}
	s.logger.Print("admin.admin user_stats")

	st1 := "2018/11"
	st2 := "2018/12"
	st3 := "2019/01"
	c := "key"
	size := "value"

	// StatsPlanetController_Bar: end_implement
	res = &admin.JeeekUserStats{
		Data: []*admin.VironDataType{
			&admin.VironDataType{Key:&st1, Value:1},
			&admin.VironDataType{Key:&st2, Value:2},
			&admin.VironDataType{Key:&st3, Value:5},
		},
		X: "key",
		Y: "value",
		Color: &c,
		Size: &size,
		Guide: &admin.VironGuideType{
			X: &admin.VironLabelType{Label: "年月"},
			Y: &admin.VironLabelType{Label: "人数"},
		},
	}

	return
}

// authtype controller(viron必須API)
func (s *adminsrvc) Authtype(ctx context.Context) (res admin.JeeekVironAuthtypeCollection, err error) {
	s.logger.Print("admin.authtype")
	res = admin.JeeekVironAuthtypeCollection{
		&admin.JeeekVironAuthtype{
			Method:   "POST",
			Provider: "",
			Type:     "jwt",
			URL:      "/v1/admin/signin",
		},
		&admin.JeeekVironAuthtype{
			Method:   "POST",
			Provider: "",
			Type:     "signout",
			URL:      "/vi/admin/signout",
		},
	}
	return
}

// viron_menu(viron必須API)
func (s *adminsrvc) VironMenu(ctx context.Context) (res *admin.JeeekVironMenu, err error) {
	res = &admin.JeeekVironMenu{}
	s.logger.Print("admin.viron_menu")
	th := "standard"
	cl := "green"
	pk := "id"
	pagenation := true

	res = &admin.JeeekVironMenu{
		Name: "Admin Screen of Jeeek",
		Tags: []string{
			"local",
		},
		Theme: &th,
		Color: &cl,
		Pages: []*admin.VironPage{
			&admin.VironPage{
				Section:    "dashboard",
				Name:       "ダッシュボード",
				ID:         "quickview",
				Components: []*admin.VironComponent{
					&admin.VironComponent{
						Name:    "Users(bar)",
						API: &admin.VironAPI{
							Method: "get",
							Path:   "/v1/admin/user_number",
						},
						Style: "graph-bar",
					},
				},
			},
			&admin.VironPage{
				Section: "manage",
				ID:      "user-admin",
				Name:    "ユーザ管理",
				Components: []*admin.VironComponent{
					&admin.VironComponent{
						API: &admin.VironAPI{
							Method: "get",
							Path:   "/v1/admin/users",
						},
						Name:       "ユーザ一覧",
						Style:      "table",
						Primary:    &pk,
						Pagination: &pagenation,
						Query: []*admin.VironQuery{
							&admin.VironQuery{
								Key:  "user_id",
								Type: "integer",
							},
							&admin.VironQuery{
								Key:  "user_name",
								Type: "string",
							},
							&admin.VironQuery{
								Key:  "email_address",
								Type: "string",
							},
							&admin.VironQuery{
								Key:  "phone_number",
								Type: "string",
							},
							&admin.VironQuery{
								Key:  "photo_url",
								Type: "string",
							},
							&admin.VironQuery{
								Key:  "email_verified",
								Type: "string",
							},
							&admin.VironQuery{
								Key:  "created_at",
								Type: "string",
							},
							&admin.VironQuery{
								Key:  "updated_at",
								Type: "string",
							},
						},
						TableLabels: []string{
							"id",
							"name",
							"email",
							"phone",
							"photo",
							"created_at",
							"updated_at",
						},
					},
				},
			},
		},
	}
	return
}
