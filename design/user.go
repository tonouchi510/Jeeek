package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("User", func() {
	Description("ユーザー/セッションに関するエンドポイントです。")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/v1")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("Get current user", func() {
		Description("現在のセッションに紐付くユーザの情報を返します。")

		Payload(SessionTokenPayload)
		Result(UserResponse)

		HTTP(func() {
			GET("/users/me")
			// Use Authorization header to provide basic auth value.
			Response(StatusOK)
		})
	})

	Method("Update user", func() {
		Description("現在のセッションに紐付くユーザー情報を更新します。")

		Payload(UpdateUserPayload)
		Result(UserResponse)

		HTTP(func() {
			PUT("/users")
			Response(StatusOK)
		})
	})

	Method("List user", func() {
		Description("ユーザーの一覧を返します。")

		Payload(SessionTokenPayload)
		Result(CollectionOf(UserResponse))

		HTTP(func() {
			GET("/users")
			Response(StatusOK)
		})
	})

	Method("Get user", func() {
		Description("指定したIDのユーザーの情報を返します。")

		Payload(GetUserPayload)
		Result(UserResponse)

		HTTP(func() {
			GET("/users/{user_id}")
			Response(StatusOK)
		})
	})

	Method("Delete user", func() {
		Description("現在のセッションに紐づくユーザーを削除します。")

		Payload(SessionTokenPayload)

		HTTP(func() {
			DELETE("/users")
			Response(StatusNoContent)
		})
	})
})
