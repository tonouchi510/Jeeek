package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Admin", func() {
	Description("管理者用のAPI。")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/v1/admin")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("admin health-check", func() {
		Description("admin apiのhealth-check")

		Payload(SessionTokenPayload)
		Result(HealthCheckResponse)

		HTTP(func() {
			GET("/healthcheck")
			Response(StatusOK)
		})
	})

	Method("admin signin", func() {
		Description("admin権限のトークンを取得します．")
		NoSecurity()

		Payload(AdminSignInPayload)
		Result(JWTResponse)

		HTTP(func() {
			POST("/signin")
			Response(StatusOK)
		})
	})

	Method("admin create new user", func(){
		Description("新しいユーザーを登録します。")

		Payload(AdminCreateUserPayload)
		Result(UserResponse)

		HTTP(func(){
			POST("/users")
			Response(StatusCreated)
		})
	})

	Method("admin update user", func() {
		Description("指定したユーザー情報を更新します。")

		Payload(AdminUpdateUserPayload)
		Result(UserResponse)

		HTTP(func() {
			PUT("/users/{user_id}")
			Response(StatusOK)
		})
	})

	Method("admin list user", func() {
		Description("ユーザーの一覧を返します。")

		Payload(SessionTokenPayload)
		Result(CollectionOf(UserResponse))

		HTTP(func() {
			GET("/users")
			Response(StatusOK)
		})
	})

	Method("admin get user", func() {
		Description("指定したIDのユーザーの情報を返します。")

		Payload(GetUserPayload)
		Result(UserResponse)

		HTTP(func() {
			GET("/users/{user_id}")
			Response(StatusOK)
		})
	})

	Method("admin delete user", func() {
		Description("指定したユーザーを削除します。")

		Payload(AdminDeleteUserPayload)

		HTTP(func() {
			DELETE("/users/{user_id}")
			Response(StatusNoContent)
		})
	})

	Method("admin user_stats", func() {
		Description("ユーザ数の統計情報を返す")

		Payload(SessionTokenPayload)
		Result(AdminUserStats)

		HTTP(func() {
			POST("/user_stats")
			Response(StatusOK)
		})
	})

	Method("authtype", func() {
		Description("authtype controller(viron必須API)")
		NoSecurity()

		Result(CollectionOf(VironAuthTypeResponse))
		HTTP(func() {
			GET("/viron_authtype")
			Response(StatusOK)
		})
	})

	Method("viron_menu", func() {
		Description("viron_menu(viron必須API)")
		NoSecurity()

		Result(VironMenuResponse)
		HTTP(func() {
			GET("/viron")
			Response(StatusOK)
		})
	})

})
