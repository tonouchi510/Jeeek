package design

import (
	. "goa.design/goa/v3/dsl"
)


// User
var UserResponse = ResultType("application/vnd.jeeek.user+json", func() {
	Description("user response")
	ContentType("application/json; charset=utf-8")

	Reference(UserProfile)
	Attributes(func() {
		Attribute("user_id")
		Attribute("user_name")
		Attribute("email_address")
		Attribute("phone_number")
		Attribute("photo_url")

		Attribute("email_verified", Boolean, func() {
			Description("ユーザーのプライマリ メールアドレスが確認されているかどうか")
		})

		Required("user_id", "user_name", "email_address", "photo_url", "phone_number", "email_verified")
	})

	View("default", func() {
		Attribute("user_id")
		Attribute("user_name")
		Attribute("email_address")
		Attribute("phone_number")
		Attribute("photo_url")
		Attribute("email_verified")
	})

	View("tiny", func() {
		Attribute("user_id")
		Attribute("user_name")
		Attribute("email_address")
	})
})
