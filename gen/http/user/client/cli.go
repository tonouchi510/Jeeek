// Code generated by goa v3.0.4, DO NOT EDIT.
//
// User HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package client

import (
	"encoding/json"
	"fmt"

	user "github.com/tonouchi510/Jeeek/gen/user"
)

// BuildGetCurrentUserPayload builds the payload for the User Get current user
// endpoint from CLI flags.
func BuildGetCurrentUserPayload(userGetCurrentUserToken string) (*user.SessionTokenPayload, error) {
	var token *string
	{
		if userGetCurrentUserToken != "" {
			token = &userGetCurrentUserToken
		}
	}
	payload := &user.SessionTokenPayload{
		Token: token,
	}
	return payload, nil
}

// BuildUpdateUserPayload builds the payload for the User Update user endpoint
// from CLI flags.
func BuildUpdateUserPayload(userUpdateUserBody string, userUpdateUserToken string) (*user.UpdateUserPayload, error) {
	var err error
	var body UpdateUserRequestBody
	{
		err = json.Unmarshal([]byte(userUpdateUserBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, example of valid JSON:\n%s", "'{\n      \"email_address\": \"keisuke.honda+testuser@ynu.jp\",\n      \"phone_number\": \"08079469367\",\n      \"photo_url\": \"https://imageurl.com\",\n      \"user_name\": \"keisuke.honda\"\n   }'")
		}
	}
	var token *string
	{
		if userUpdateUserToken != "" {
			token = &userUpdateUserToken
		}
	}
	v := &user.UpdateUserPayload{
		UserName:     body.UserName,
		EmailAddress: body.EmailAddress,
		PhoneNumber:  body.PhoneNumber,
		PhotoURL:     body.PhotoURL,
	}
	v.Token = token
	return v, nil
}

// BuildListUserPayload builds the payload for the User List user endpoint from
// CLI flags.
func BuildListUserPayload(userListUserToken string) (*user.SessionTokenPayload, error) {
	var token *string
	{
		if userListUserToken != "" {
			token = &userListUserToken
		}
	}
	payload := &user.SessionTokenPayload{
		Token: token,
	}
	return payload, nil
}

// BuildGetUserPayload builds the payload for the User Get user endpoint from
// CLI flags.
func BuildGetUserPayload(userGetUserUserID string, userGetUserToken string) (*user.GetUserPayload, error) {
	var userID string
	{
		userID = userGetUserUserID
	}
	var token *string
	{
		if userGetUserToken != "" {
			token = &userGetUserToken
		}
	}
	payload := &user.GetUserPayload{
		UserID: userID,
		Token:  token,
	}
	return payload, nil
}

// BuildDeleteUserPayload builds the payload for the User Delete user endpoint
// from CLI flags.
func BuildDeleteUserPayload(userDeleteUserToken string) (*user.SessionTokenPayload, error) {
	var token *string
	{
		if userDeleteUserToken != "" {
			token = &userDeleteUserToken
		}
	}
	payload := &user.SessionTokenPayload{
		Token: token,
	}
	return payload, nil
}
