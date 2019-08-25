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
