// Code generated by goa v3.0.4, DO NOT EDIT.
//
// Admin views
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package views

import (
	"unicode/utf8"

	goa "goa.design/goa/v3/pkg"
)

// JeeekHealthcheck is the viewed result type that is projected based on a view.
type JeeekHealthcheck struct {
	// Type to project
	Projected *JeeekHealthcheckView
	// View to render
	View string
}

// JeeekAdminSignin is the viewed result type that is projected based on a view.
type JeeekAdminSignin struct {
	// Type to project
	Projected *JeeekAdminSigninView
	// View to render
	View string
}

// JeeekUser is the viewed result type that is projected based on a view.
type JeeekUser struct {
	// Type to project
	Projected *JeeekUserView
	// View to render
	View string
}

// JeeekUserCollection is the viewed result type that is projected based on a
// view.
type JeeekUserCollection struct {
	// Type to project
	Projected JeeekUserCollectionView
	// View to render
	View string
}

// JeeekHealthcheckView is a type that runs validations on a projected type.
type JeeekHealthcheckView struct {
	Result *string
}

// JeeekAdminSigninView is a type that runs validations on a projected type.
type JeeekAdminSigninView struct {
	// JWT used for authentication
	Token *string
}

// JeeekUserView is a type that runs validations on a projected type.
type JeeekUserView struct {
	// User id of firebase
	UserID *string
	// ユーザーの表示名
	UserName *string
	// ユーザーのプライマリ メールアドレス
	EmailAddress *string
	// ユーザのメインの電話番号
	PhoneNumber *string
	// ユーザーの写真 URL
	PhotoURL *string
	// ユーザーのプライマリ メールアドレスが確認されているかどうか
	EmailVerified *bool
	// ユーザが停止状態かどうか（論理削除）
	Disabled *bool
	// ユーザが作成された日時
	CreatedAt *string
	// 最後にログインした日時
	LastSignedinAt *string
}

// JeeekUserCollectionView is a type that runs validations on a projected type.
type JeeekUserCollectionView []*JeeekUserView

var (
	// JeeekHealthcheckMap is a map of attribute names in result type
	// JeeekHealthcheck indexed by view name.
	JeeekHealthcheckMap = map[string][]string{
		"default": []string{
			"result",
		},
	}
	// JeeekAdminSigninMap is a map of attribute names in result type
	// JeeekAdminSignin indexed by view name.
	JeeekAdminSigninMap = map[string][]string{
		"default": []string{
			"token",
		},
	}
	// JeeekUserMap is a map of attribute names in result type JeeekUser indexed by
	// view name.
	JeeekUserMap = map[string][]string{
		"default": []string{
			"user_id",
			"user_name",
			"email_address",
			"phone_number",
			"photo_url",
			"email_verified",
		},
		"tiny": []string{
			"user_id",
			"user_name",
			"email_address",
		},
		"admin": []string{
			"user_id",
			"user_name",
			"email_address",
			"phone_number",
			"photo_url",
			"email_verified",
			"disabled",
			"created_at",
			"last_signedin_at",
		},
	}
	// JeeekUserCollectionMap is a map of attribute names in result type
	// JeeekUserCollection indexed by view name.
	JeeekUserCollectionMap = map[string][]string{
		"default": []string{
			"user_id",
			"user_name",
			"email_address",
			"phone_number",
			"photo_url",
			"email_verified",
		},
		"tiny": []string{
			"user_id",
			"user_name",
			"email_address",
		},
		"admin": []string{
			"user_id",
			"user_name",
			"email_address",
			"phone_number",
			"photo_url",
			"email_verified",
			"disabled",
			"created_at",
			"last_signedin_at",
		},
	}
)

// ValidateJeeekHealthcheck runs the validations defined on the viewed result
// type JeeekHealthcheck.
func ValidateJeeekHealthcheck(result *JeeekHealthcheck) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJeeekHealthcheckView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateJeeekAdminSignin runs the validations defined on the viewed result
// type JeeekAdminSignin.
func ValidateJeeekAdminSignin(result *JeeekAdminSignin) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJeeekAdminSigninView(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default"})
	}
	return
}

// ValidateJeeekUser runs the validations defined on the viewed result type
// JeeekUser.
func ValidateJeeekUser(result *JeeekUser) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJeeekUserView(result.Projected)
	case "tiny":
		err = ValidateJeeekUserViewTiny(result.Projected)
	case "admin":
		err = ValidateJeeekUserViewAdmin(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny", "admin"})
	}
	return
}

// ValidateJeeekUserCollection runs the validations defined on the viewed
// result type JeeekUserCollection.
func ValidateJeeekUserCollection(result JeeekUserCollection) (err error) {
	switch result.View {
	case "default", "":
		err = ValidateJeeekUserCollectionView(result.Projected)
	case "tiny":
		err = ValidateJeeekUserCollectionViewTiny(result.Projected)
	case "admin":
		err = ValidateJeeekUserCollectionViewAdmin(result.Projected)
	default:
		err = goa.InvalidEnumValueError("view", result.View, []interface{}{"default", "tiny", "admin"})
	}
	return
}

// ValidateJeeekHealthcheckView runs the validations defined on
// JeeekHealthcheckView using the "default" view.
func ValidateJeeekHealthcheckView(result *JeeekHealthcheckView) (err error) {
	if result.Result == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("result", "result"))
	}
	return
}

// ValidateJeeekAdminSigninView runs the validations defined on
// JeeekAdminSigninView using the "default" view.
func ValidateJeeekAdminSigninView(result *JeeekAdminSigninView) (err error) {
	if result.Token == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("token", "result"))
	}
	return
}

// ValidateJeeekUserView runs the validations defined on JeeekUserView using
// the "default" view.
func ValidateJeeekUserView(result *JeeekUserView) (err error) {
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "result"))
	}
	if result.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "result"))
	}
	if result.EmailAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email_address", "result"))
	}
	if result.PhotoURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("photo_url", "result"))
	}
	if result.PhoneNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("phone_number", "result"))
	}
	if result.EmailVerified == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email_verified", "result"))
	}
	if result.UserID != nil {
		if utf8.RuneCountInString(*result.UserID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_id", *result.UserID, utf8.RuneCountInString(*result.UserID), 28, true))
		}
	}
	if result.UserID != nil {
		if utf8.RuneCountInString(*result.UserID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_id", *result.UserID, utf8.RuneCountInString(*result.UserID), 28, false))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 1, true))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 20, false))
		}
	}
	if result.EmailAddress != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.email_address", *result.EmailAddress, goa.FormatEmail))
	}
	if result.PhoneNumber != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.phone_number", *result.PhoneNumber, "^\\+?[\\d]{10,}$"))
	}
	return
}

// ValidateJeeekUserViewTiny runs the validations defined on JeeekUserView
// using the "tiny" view.
func ValidateJeeekUserViewTiny(result *JeeekUserView) (err error) {
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "result"))
	}
	if result.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "result"))
	}
	if result.EmailAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email_address", "result"))
	}
	if result.UserID != nil {
		if utf8.RuneCountInString(*result.UserID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_id", *result.UserID, utf8.RuneCountInString(*result.UserID), 28, true))
		}
	}
	if result.UserID != nil {
		if utf8.RuneCountInString(*result.UserID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_id", *result.UserID, utf8.RuneCountInString(*result.UserID), 28, false))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 1, true))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 20, false))
		}
	}
	if result.EmailAddress != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.email_address", *result.EmailAddress, goa.FormatEmail))
	}
	return
}

// ValidateJeeekUserViewAdmin runs the validations defined on JeeekUserView
// using the "admin" view.
func ValidateJeeekUserViewAdmin(result *JeeekUserView) (err error) {
	if result.UserID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_id", "result"))
	}
	if result.UserName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("user_name", "result"))
	}
	if result.EmailAddress == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email_address", "result"))
	}
	if result.PhotoURL == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("photo_url", "result"))
	}
	if result.PhoneNumber == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("phone_number", "result"))
	}
	if result.EmailVerified == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email_verified", "result"))
	}
	if result.UserID != nil {
		if utf8.RuneCountInString(*result.UserID) < 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_id", *result.UserID, utf8.RuneCountInString(*result.UserID), 28, true))
		}
	}
	if result.UserID != nil {
		if utf8.RuneCountInString(*result.UserID) > 28 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_id", *result.UserID, utf8.RuneCountInString(*result.UserID), 28, false))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 1, true))
		}
	}
	if result.UserName != nil {
		if utf8.RuneCountInString(*result.UserName) > 20 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("result.user_name", *result.UserName, utf8.RuneCountInString(*result.UserName), 20, false))
		}
	}
	if result.EmailAddress != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("result.email_address", *result.EmailAddress, goa.FormatEmail))
	}
	if result.PhoneNumber != nil {
		err = goa.MergeErrors(err, goa.ValidatePattern("result.phone_number", *result.PhoneNumber, "^\\+?[\\d]{10,}$"))
	}
	return
}

// ValidateJeeekUserCollectionView runs the validations defined on
// JeeekUserCollectionView using the "default" view.
func ValidateJeeekUserCollectionView(result JeeekUserCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateJeeekUserView(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateJeeekUserCollectionViewTiny runs the validations defined on
// JeeekUserCollectionView using the "tiny" view.
func ValidateJeeekUserCollectionViewTiny(result JeeekUserCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateJeeekUserViewTiny(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateJeeekUserCollectionViewAdmin runs the validations defined on
// JeeekUserCollectionView using the "admin" view.
func ValidateJeeekUserCollectionViewAdmin(result JeeekUserCollectionView) (err error) {
	for _, item := range result {
		if err2 := ValidateJeeekUserViewAdmin(item); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}
