package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Activity", func() {
	Description("アクティビティ投稿関連のAPI")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/v1/activity")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("Manual post of activity", func() {
		Description("手動投稿用のAPI")

		Payload(ActivityPostPayload)

		HTTP(func() {
			POST("/post")
			Response(StatusOK)
		})
	})

	Method("Refresh activities of all cooperation services", func() {
		Description("セッションに紐づくユーザの連携済みサービスのアクティビティを取得する")

		Payload(SessionTokenPayload)

		HTTP(func() {
			GET("/co-service/batch")
			Response(StatusOK)
		})
	})

	Method("Refresh qiita activities", func() {
		Description("セッションに紐づくユーザのQiitaの記事投稿を取得する")

		Payload(SessionTokenPayload)

		HTTP(func() {
			GET("/co-service/qiita")
			Response(StatusOK)
		})
	})

	Method("Pick out all past activities of qiita", func() {
		Description("サービス連携時に連携以前のqiita記事投稿を全て反映させる")

		Payload(SessionTokenPayload)

		HTTP(func() {
			GET("/co-service/qiita/initialization")
			Response(StatusOK)
		})
	})

})
