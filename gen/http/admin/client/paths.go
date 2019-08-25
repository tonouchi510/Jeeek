// Code generated by goa v3.0.4, DO NOT EDIT.
//
// HTTP request path constructors for the Admin service.
//
// Command:
// $ goa gen github.com/tonouchi510/Jeeek/design

package client

import (
	"fmt"
)

// AdminHealthCheckAdminPath returns the URL path to the Admin service admin health-check HTTP endpoint.
func AdminHealthCheckAdminPath() string {
	return "/v1/admin/healthcheck"
}

// AdminSigninAdminPath returns the URL path to the Admin service admin signin HTTP endpoint.
func AdminSigninAdminPath() string {
	return "/v1/admin/signin"
}

// AdminCreateNewUserAdminPath returns the URL path to the Admin service admin create new user HTTP endpoint.
func AdminCreateNewUserAdminPath() string {
	return "/v1/admin/users"
}

// AdminUpdateUserAdminPath returns the URL path to the Admin service admin update user HTTP endpoint.
func AdminUpdateUserAdminPath(userID string) string {
	return fmt.Sprintf("/v1/admin/users/%v", userID)
}

// AdminListUserAdminPath returns the URL path to the Admin service admin list user HTTP endpoint.
func AdminListUserAdminPath() string {
	return "/v1/admin/users"
}

// AdminGetUserAdminPath returns the URL path to the Admin service admin get user HTTP endpoint.
func AdminGetUserAdminPath(userID string) string {
	return fmt.Sprintf("/v1/admin/users/%v", userID)
}

// AdminDeleteUserAdminPath returns the URL path to the Admin service admin delete user HTTP endpoint.
func AdminDeleteUserAdminPath(userID string) string {
	return fmt.Sprintf("/v1/admin/users/%v", userID)
}

// AdminUserStatsAdminPath returns the URL path to the Admin service admin user_stats HTTP endpoint.
func AdminUserStatsAdminPath() string {
	return "/v1/admin/user_stats"
}

// AuthtypeAdminPath returns the URL path to the Admin service authtype HTTP endpoint.
func AuthtypeAdminPath() string {
	return "/v1/admin/viron_authtype"
}

// VironMenuAdminPath returns the URL path to the Admin service viron_menu HTTP endpoint.
func VironMenuAdminPath() string {
	return "/v1/admin/viron"
}
