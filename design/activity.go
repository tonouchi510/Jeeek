package design

import (
	. "goa.design/goa/v3/dsl"
)

var _ = Service("Activity", func() {
	Description("フォロワーへのアクティビティ投稿の反映API")

	Security(JWTAuth)
	Error("unauthorized", String, "Credentials are invalid")

	HTTP(func() {
		Path("/v1/activity")
		Response("unauthorized", StatusUnauthorized)
	})

	Method("Reflection activity", func() {
		Description("タイムラインへの書き込みを行う")

		Payload(ActivityPostPayload)

		HTTP(func() {
			POST("/post")
			Response(StatusOK)
		})
	})

})
