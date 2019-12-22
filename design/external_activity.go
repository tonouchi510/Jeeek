package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("ExternalActivity", func() {
	Description("外部サービスアクティビティの自動取得API")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/v1/external_activity")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("Refresh activities of external services", func() {
		Description("セッションに紐づくユーザの連携済みサービスのアクティビティを取得する")

		Payload(SessionTokenPayload)

		HTTP(func() {
			GET("/batch")
			Response(StatusOK)
		})
	})

	Method("Refresh qiita activity", func() {
		Description("セッションに紐づくユーザのQiitaの記事投稿を取得する")

		Payload(SessionTokenPayload)

		HTTP(func() {
			GET("/qiita")
			Response(StatusOK)
		})
	})

	Method("Pick out past activity of qiita", func() {
		Description("サービス連携時に連携以前のqiita記事投稿を全て反映させる")

		Payload(SessionTokenPayload)

		HTTP(func() {
			GET("/qiita/initialization")
			Response(StatusOK)
		})
	})

})
