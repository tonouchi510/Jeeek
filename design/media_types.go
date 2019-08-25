package design

import (
	. "goa.design/goa/v3/dsl"
)

var HealthCheckResponse = ResultType("application/vnd.jeeek.healthcheck+json", func() {
	Description("health-check response")
	ContentType("application/json; charset=utf-8")

	Attributes(func() {
		Attribute("result", String, func() {
			Default("OK")
			Example("OK")
		})
		Required("result")
	})
})


//Admin
var JWTResponse = ResultType("application/vnd.jeeek.admin.signin+json", func() {
	Description("admin-signin response")
	ContentType("application/json; charset=utf-8")

	Reference(JWT)
	Attributes(func() {
		Attribute("token")
		Required("token")
	})
})

// MediaType of Viron API
// for /viron_authtype
var VironAuthTypeResponse = ResultType("application/vnd.jeeek.viron.authtype+json", func() {
	Attributes(func() {

		Attribute("type", String, "type name")
		Attribute("provider", String, "provider name")
		Attribute("url", String, "url")
		Attribute("method", String, "request method to submit this auth")

		Required("type", "provider", "url", "method")
	})

	View("default", func() {
		Attribute("type")
		Attribute("provider")
		Attribute("url")
		Attribute("method")
	})
})

// MenuType menu information on /viron
var VironMenuResponse = ResultType("application/vnd.jeeek.viron.menu+json", func() {
	Attributes(func() {
		Attribute("theme", String, func() {
			Enum("standard", "midnight", "terminal")
		})
		Attribute("color", String, func() {
			Enum("purple", "blue", "green", "yellow", "red", "gray", "black", "white")
		})
		Attribute("name", String)
		Attribute("tags", ArrayOf(String))
		Attribute("thumbnail", String)
		Attribute("pages", ArrayOf(VironPage))
		Attribute("sections", ArrayOf(VironSection))

		Required("name", "pages")
	})
	View("default", func() {
		Attribute("theme")
		Attribute("color")
		Attribute("name")
		Attribute("tags")
		Attribute("thumbnail")
		Attribute("pages")
	})
})

var AdminUserStats = ResultType("application/vnd.jeeek.user.stats+json", func() {
	Description("user-stats response")
	ContentType("application/json; charset=utf-8")

	Attributes(func() {
		Attribute("data", ArrayOf(VironDataType), "グラフデータ")
		Attribute("x", String, "X軸に使用するkey")
		Attribute("y", String, "Y軸に使用するkey")
		Attribute("size", String, "ドットの大きさに使用するkey")
		Attribute("color", String, "ドットの色分けに使用するkey")
		Attribute("guide", VironGuideType)

		Required("data", "x", "y", "guide")
	})
})


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
		Attribute("email_verified")
		Attribute("disabled")
		Attribute("created_at")
		Attribute("last_signedin_at")

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

	View("admin", func() {
		Attribute("user_id")
		Attribute("user_name")
		Attribute("email_address")
		Attribute("phone_number")
		Attribute("photo_url")
		Attribute("email_verified")
		Attribute("disabled")
		Attribute("created_at")
		Attribute("last_signedin_at")
	})

})
