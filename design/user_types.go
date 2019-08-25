package design

import (
	. "goa.design/goa/v3/dsl"
)

var digitPattern = `^[0-9]+$`
var phoneNumberPattern = `^\+?[\d]{10,}$`

var IDKeyDefinition = func() {
	Description("IDKey of datastore")
	Example(5644004762845184)
}

var JWT = Type("JWT Token", func() {
	Token("token", String, func() {
		Description("JWT used for authentication")
		Example("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ")
	})
})

var SessionTokenPayload = Type("SessionTokenPayload", func() {
	Reference(JWT)
	Token("token")
})


var TimeStamp = Type("TimeStamp", func() {
	Attribute("created_at", String, func() {
		Example("yyyy-mm-dd:xx:yy:zz")
	})
	Attribute("updated_at", String, func() {
		Example("yyyy-mm-dd:xx:yy:zz")
	})
	Attribute("deleted_at", String, func() {
		Example("yyyy-mm-dd:xx:yy:zz")
	})
})


//Admin
var AdminSignInPayload = Type("AdminSignInPayload", func() {
	Attribute("uid", String, func() {
		Example("4WIbqiNIpIYXqrfBMVZsbKCepau1")
	})
	Required("uid")
})

// VironSection component for section in viron
var VironSection = Type("VironSection", func() {
	Attribute("id", String)
	Attribute("label", String)

	Required("id", "label")
})

// VironPage pagetype media
var VironPage = Type("VironPage", func() {
	Attribute("section", String, func() {
		Description("中カテゴリのセクション")
	})
	Attribute("group", String, func() {
		Default("")
	})
	Attribute("id", String)
	Attribute("name", String)
	Attribute("components", ArrayOf(VironComponent))

	Required("section", "group", "id", "name", "components")
})

// VironComponent media type for component in viron
var VironComponent = Type("VironComponent", func() {
	Attribute("api", VironAPI)
	Attribute("name", String)
	Attribute("style", String, func() {
		Enum("number",
			"table",
			"graph-bar",
			"graph-scatterplot",
			"graph-line",
			"graph-horizontal-bar",
			"graph-stacked-bar",
			"graph-horizontal-stacked-bar",
			"graph-stacked-area",
		)
	})
	Attribute("auto_refresh_sec", Int32, func() {
		Description("指定された秒数毎に自動でデータを更新")
	})
	Attribute("primary", String)
	Attribute("pagination", Boolean)
	Attribute("query", ArrayOf(VironQuery))
	Attribute("table_labels", ArrayOf(String))
	Attribute("actions", ArrayOf(String), func() {
		Description("指定フォーマットに合わないURIのAPIを追加")
	})
	Attribute("unit", String)

	Required("name", "style", "api")
})

// VironAPI media type for api in viron
var VironAPI = Type("VironAPI", func() {
	Attribute("method", String)
	Attribute("path", String)

	Required("method", "path")
})

// VironQuery mediatype for query in viron
var VironQuery = Type("VironQuery", func() {
	Attribute("key", String)
	Attribute("type", String)

	Required("key", "type")
})


// User
var UserProfile = Type("UserProfile", func() {
	Attribute("user_id", String, UserIDDefinition)

	Attribute("user_name", String, func() {
		Description("ユーザーの表示名")
		MinLength(1)
		MaxLength(20)
		Example("keisuke.honda")
	})

	Attribute("email_address", String, func() {
		Description("ーザーのプライマリ メールアドレス")
		Format("email")
		Example("keisuke.honda+testuser@ynu.jp")
	})

	Attribute("phone_number", String, func() {
		Description("ユーザのメインの電話番号")
		Pattern(phoneNumberPattern)
		Example("08079469367")
	})

	Attribute("photo_url", String, func() {
		Description("ユーザーの写真 URL")
		Example("https://imageurl.com")
	})
})

var UserIDDefinition = func() {
	Description("User id of firebase")
	Example("XRQ85mtXnINISH25zfM0m5RlC6L2")
	MinLength(28)
	MaxLength(28)
}

var GetUserPayload = Type("GetUserPayload", func() {
	Reference(JWT)
	Token("token")
	Attribute("user_id", String, UserIDDefinition)
	Required("user_id")
})

var UpdateUserPayload = Type("UpdateUserPayload", func() {
	Reference(JWT)
	Token("token")
	Reference(UserProfile)
	Attribute("user_name")
	Attribute("email_address")
	Attribute("phone_number")
	Attribute("photo_url")
})
