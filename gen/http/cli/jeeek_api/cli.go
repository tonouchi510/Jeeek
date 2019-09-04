// Code generated by goa v3.0.4, DO NOT EDIT.
//
// JeeekAPI HTTP client CLI support package
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	adminc "github.com/tonouchi510/Jeeek/gen/http/admin/client"
	userc "github.com/tonouchi510/Jeeek/gen/http/user/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `admin (admin- health--check|admin- signin|admin- create- new- user|admin- update- user|admin- list- user|admin- get- user|admin- delete- user)
user (get- current- user|update- user|list- user|get- user|delete- user)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` admin admin--- health----check --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"` + "\n" +
		os.Args[0] + ` user get--- current--- user --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		adminFlags = flag.NewFlagSet("admin", flag.ContinueOnError)

		adminAdminHealthCheckFlags     = flag.NewFlagSet("admin- health--check", flag.ExitOnError)
		adminAdminHealthCheckTokenFlag = adminAdminHealthCheckFlags.String("token", "", "")

		adminAdminSigninFlags    = flag.NewFlagSet("admin- signin", flag.ExitOnError)
		adminAdminSigninBodyFlag = adminAdminSigninFlags.String("body", "REQUIRED", "")

		adminAdminCreateNewUserFlags     = flag.NewFlagSet("admin- create- new- user", flag.ExitOnError)
		adminAdminCreateNewUserBodyFlag  = adminAdminCreateNewUserFlags.String("body", "REQUIRED", "")
		adminAdminCreateNewUserTokenFlag = adminAdminCreateNewUserFlags.String("token", "", "")

		adminAdminUpdateUserFlags      = flag.NewFlagSet("admin- update- user", flag.ExitOnError)
		adminAdminUpdateUserBodyFlag   = adminAdminUpdateUserFlags.String("body", "REQUIRED", "")
		adminAdminUpdateUserUserIDFlag = adminAdminUpdateUserFlags.String("user-id", "REQUIRED", "User id of firebase")
		adminAdminUpdateUserTokenFlag  = adminAdminUpdateUserFlags.String("token", "", "")

		adminAdminListUserFlags     = flag.NewFlagSet("admin- list- user", flag.ExitOnError)
		adminAdminListUserTokenFlag = adminAdminListUserFlags.String("token", "", "")

		adminAdminGetUserFlags      = flag.NewFlagSet("admin- get- user", flag.ExitOnError)
		adminAdminGetUserUserIDFlag = adminAdminGetUserFlags.String("user-id", "REQUIRED", "User id of firebase")
		adminAdminGetUserTokenFlag  = adminAdminGetUserFlags.String("token", "", "")

		adminAdminDeleteUserFlags      = flag.NewFlagSet("admin- delete- user", flag.ExitOnError)
		adminAdminDeleteUserUserIDFlag = adminAdminDeleteUserFlags.String("user-id", "REQUIRED", "User id of firebase")
		adminAdminDeleteUserTokenFlag  = adminAdminDeleteUserFlags.String("token", "", "")

		userFlags = flag.NewFlagSet("user", flag.ContinueOnError)

		userGetCurrentUserFlags     = flag.NewFlagSet("get- current- user", flag.ExitOnError)
		userGetCurrentUserTokenFlag = userGetCurrentUserFlags.String("token", "", "")

		userUpdateUserFlags     = flag.NewFlagSet("update- user", flag.ExitOnError)
		userUpdateUserBodyFlag  = userUpdateUserFlags.String("body", "REQUIRED", "")
		userUpdateUserTokenFlag = userUpdateUserFlags.String("token", "", "")

		userListUserFlags     = flag.NewFlagSet("list- user", flag.ExitOnError)
		userListUserTokenFlag = userListUserFlags.String("token", "", "")

		userGetUserFlags      = flag.NewFlagSet("get- user", flag.ExitOnError)
		userGetUserUserIDFlag = userGetUserFlags.String("user-id", "REQUIRED", "User id of firebase")
		userGetUserTokenFlag  = userGetUserFlags.String("token", "", "")

		userDeleteUserFlags     = flag.NewFlagSet("delete- user", flag.ExitOnError)
		userDeleteUserTokenFlag = userDeleteUserFlags.String("token", "", "")
	)
	adminFlags.Usage = adminUsage
	adminAdminHealthCheckFlags.Usage = adminAdminHealthCheckUsage
	adminAdminSigninFlags.Usage = adminAdminSigninUsage
	adminAdminCreateNewUserFlags.Usage = adminAdminCreateNewUserUsage
	adminAdminUpdateUserFlags.Usage = adminAdminUpdateUserUsage
	adminAdminListUserFlags.Usage = adminAdminListUserUsage
	adminAdminGetUserFlags.Usage = adminAdminGetUserUsage
	adminAdminDeleteUserFlags.Usage = adminAdminDeleteUserUsage

	userFlags.Usage = userUsage
	userGetCurrentUserFlags.Usage = userGetCurrentUserUsage
	userUpdateUserFlags.Usage = userUpdateUserUsage
	userListUserFlags.Usage = userListUserUsage
	userGetUserFlags.Usage = userGetUserUsage
	userDeleteUserFlags.Usage = userDeleteUserUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "admin":
			svcf = adminFlags
		case "user":
			svcf = userFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "admin":
			switch epn {
			case "admin- health--check":
				epf = adminAdminHealthCheckFlags

			case "admin- signin":
				epf = adminAdminSigninFlags

			case "admin- create- new- user":
				epf = adminAdminCreateNewUserFlags

			case "admin- update- user":
				epf = adminAdminUpdateUserFlags

			case "admin- list- user":
				epf = adminAdminListUserFlags

			case "admin- get- user":
				epf = adminAdminGetUserFlags

			case "admin- delete- user":
				epf = adminAdminDeleteUserFlags

			}

		case "user":
			switch epn {
			case "get- current- user":
				epf = userGetCurrentUserFlags

			case "update- user":
				epf = userUpdateUserFlags

			case "list- user":
				epf = userListUserFlags

			case "get- user":
				epf = userGetUserFlags

			case "delete- user":
				epf = userDeleteUserFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "admin":
			c := adminc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "admin- health--check":
				endpoint = c.AdminHealthCheck()
				data, err = adminc.BuildAdminHealthCheckPayload(*adminAdminHealthCheckTokenFlag)
			case "admin- signin":
				endpoint = c.AdminSignin()
				data, err = adminc.BuildAdminSigninPayload(*adminAdminSigninBodyFlag)
			case "admin- create- new- user":
				endpoint = c.AdminCreateNewUser()
				data, err = adminc.BuildAdminCreateNewUserPayload(*adminAdminCreateNewUserBodyFlag, *adminAdminCreateNewUserTokenFlag)
			case "admin- update- user":
				endpoint = c.AdminUpdateUser()
				data, err = adminc.BuildAdminUpdateUserPayload(*adminAdminUpdateUserBodyFlag, *adminAdminUpdateUserUserIDFlag, *adminAdminUpdateUserTokenFlag)
			case "admin- list- user":
				endpoint = c.AdminListUser()
				data, err = adminc.BuildAdminListUserPayload(*adminAdminListUserTokenFlag)
			case "admin- get- user":
				endpoint = c.AdminGetUser()
				data, err = adminc.BuildAdminGetUserPayload(*adminAdminGetUserUserIDFlag, *adminAdminGetUserTokenFlag)
			case "admin- delete- user":
				endpoint = c.AdminDeleteUser()
				data, err = adminc.BuildAdminDeleteUserPayload(*adminAdminDeleteUserUserIDFlag, *adminAdminDeleteUserTokenFlag)
			}
		case "user":
			c := userc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get- current- user":
				endpoint = c.GetCurrentUser()
				data, err = userc.BuildGetCurrentUserPayload(*userGetCurrentUserTokenFlag)
			case "update- user":
				endpoint = c.UpdateUser()
				data, err = userc.BuildUpdateUserPayload(*userUpdateUserBodyFlag, *userUpdateUserTokenFlag)
			case "list- user":
				endpoint = c.ListUser()
				data, err = userc.BuildListUserPayload(*userListUserTokenFlag)
			case "get- user":
				endpoint = c.GetUser()
				data, err = userc.BuildGetUserPayload(*userGetUserUserIDFlag, *userGetUserTokenFlag)
			case "delete- user":
				endpoint = c.DeleteUser()
				data, err = userc.BuildDeleteUserPayload(*userDeleteUserTokenFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// adminUsage displays the usage of the admin command and its subcommands.
func adminUsage() {
	fmt.Fprintf(os.Stderr, `管理者用のAPI。
Usage:
    %s [globalflags] admin COMMAND [flags]

COMMAND:
    admin- health--check: admin apiのhealth-check
    admin- signin: admin権限のトークンを取得します．
    admin- create- new- user: 新しいユーザーを登録します。
    admin- update- user: 指定したユーザー情報を更新します。
    admin- list- user: ユーザーの一覧を返します。
    admin- get- user: 指定したIDのユーザーの情報を返します。
    admin- delete- user: 指定したユーザーを削除します。

Additional help:
    %s admin COMMAND --help
`, os.Args[0], os.Args[0])
}
func adminAdminHealthCheckUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- health--check -token STRING

admin apiのhealth-check
    -token STRING: 

Example:
    `+os.Args[0]+` admin admin--- health----check --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func adminAdminSigninUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- signin -body JSON

admin権限のトークンを取得します．
    -body JSON: 

Example:
    `+os.Args[0]+` admin admin--- signin --body '{
      "uid": "4WIbqiNIpIYXqrfBMVZsbKCepau1"
   }'
`, os.Args[0])
}

func adminAdminCreateNewUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- create- new- user -body JSON -token STRING

新しいユーザーを登録します。
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` admin admin--- create--- new--- user --body '{
      "email_address": "keisuke.honda+testuser@ynu.jp",
      "phone_number": "08079469367",
      "photo_url": "https://imageurl.com",
      "user_name": "keisuke.honda"
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func adminAdminUpdateUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- update- user -body JSON -user-id STRING -token STRING

指定したユーザー情報を更新します。
    -body JSON: 
    -user-id STRING: User id of firebase
    -token STRING: 

Example:
    `+os.Args[0]+` admin admin--- update--- user --body '{
      "disabled": true,
      "email_address": "keisuke.honda+testuser@ynu.jp",
      "email_verified": true,
      "phone_number": "08079469367",
      "photo_url": "https://imageurl.com",
      "user_name": "keisuke.honda"
   }' --user-id "XRQ85mtXnINISH25zfM0m5RlC6L2" --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func adminAdminListUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- list- user -token STRING

ユーザーの一覧を返します。
    -token STRING: 

Example:
    `+os.Args[0]+` admin admin--- list--- user --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func adminAdminGetUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- get- user -user-id STRING -token STRING

指定したIDのユーザーの情報を返します。
    -user-id STRING: User id of firebase
    -token STRING: 

Example:
    `+os.Args[0]+` admin admin--- get--- user --user-id "XRQ85mtXnINISH25zfM0m5RlC6L2" --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func adminAdminDeleteUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] admin admin- delete- user -user-id STRING -token STRING

指定したユーザーを削除します。
    -user-id STRING: User id of firebase
    -token STRING: 

Example:
    `+os.Args[0]+` admin admin--- delete--- user --user-id "XRQ85mtXnINISH25zfM0m5RlC6L2" --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

// userUsage displays the usage of the user command and its subcommands.
func userUsage() {
	fmt.Fprintf(os.Stderr, `ユーザー/セッションに関するエンドポイントです。
Usage:
    %s [globalflags] user COMMAND [flags]

COMMAND:
    get- current- user: 現在のセッションに紐付くユーザの情報を返します。
    update- user: 現在のセッションに紐付くユーザー情報を更新します。
    list- user: ユーザーの一覧を返します。
    get- user: 指定したIDのユーザーの情報を返します。
    delete- user: 現在のセッションに紐づくユーザーを削除します。

Additional help:
    %s user COMMAND --help
`, os.Args[0], os.Args[0])
}
func userGetCurrentUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user get- current- user -token STRING

現在のセッションに紐付くユーザの情報を返します。
    -token STRING: 

Example:
    `+os.Args[0]+` user get--- current--- user --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func userUpdateUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user update- user -body JSON -token STRING

現在のセッションに紐付くユーザー情報を更新します。
    -body JSON: 
    -token STRING: 

Example:
    `+os.Args[0]+` user update--- user --body '{
      "email_address": "keisuke.honda+testuser@ynu.jp",
      "phone_number": "08079469367",
      "photo_url": "https://imageurl.com",
      "user_name": "keisuke.honda"
   }' --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func userListUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user list- user -token STRING

ユーザーの一覧を返します。
    -token STRING: 

Example:
    `+os.Args[0]+` user list--- user --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func userGetUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user get- user -user-id STRING -token STRING

指定したIDのユーザーの情報を返します。
    -user-id STRING: User id of firebase
    -token STRING: 

Example:
    `+os.Args[0]+` user get--- user --user-id "XRQ85mtXnINISH25zfM0m5RlC6L2" --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}

func userDeleteUserUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] user delete- user -token STRING

現在のセッションに紐づくユーザーを削除します。
    -token STRING: 

Example:
    `+os.Args[0]+` user delete--- user --token "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ"
`, os.Args[0])
}
