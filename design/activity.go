package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Activity", func() {
	Description("アクティビティの自動取得API")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/v1/activity")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("Fetch qiita article by qiita-user-id", func() {
		Description("指定したユーザのQiitaの記事投稿を取得する")

		Payload(GetActivityPayload)

		HTTP(func() {
			GET("/qiita/{user_id}")
			Response(StatusOK)
		})
	})

	Method("Batch job method to refresh qiita activity", func() {
		Description("qiita連携済みユーザのqiitaでのアクティビティ更新を行うジョブ")
		NoSecurity()

		HTTP(func() {
			GET("/qiita/batch")
			Response(StatusOK)
		})
	})

	Method("Pick out past activity of qiita", func() {
		Description("サービス連携以前のqiita記事投稿を反映させる")

		Payload(GetActivityPayload)

		HTTP(func() {
			GET("/qiita/initialization")
			Response(StatusOK)
		})
	})

})
