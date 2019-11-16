package controller

import (
	"context"
	"github.com/tonouchi510/Jeeek/gen/admin"
	"github.com/tonouchi510/Jeeek/gen/user"
	"log"

	"goa.design/goa/v3/security"
)

// JWTAuth implements the authorization logic for service "Admin" for the "jwt"
// security scheme.
func (s *adminsrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	verifiedToken, err := s.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return ctx, admin.Unauthorized("invalid token")
	}

	claims := verifiedToken.Claims
	if _, ok := claims["admin"]; !ok {
		return ctx, admin.Unauthorized("You do not have administrator privileges.")
	}

	log.Printf("Verified ID token: %v\n", verifiedToken)
	return ctx, nil
}

// JWTAuth implements the authorization logic for service "User" for the "jwt"
// security scheme.
func (s *usersrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	verifiedToken, err := s.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return ctx, user.Unauthorized("invalid token")
	}

	log.Printf("Verified ID token: %v\n", verifiedToken)

	return ctx, nil
}

// JWTAuth implements the authorization logic for service "Activity" for the "jwt"
// security scheme.
func (s *activitysrvc) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	verifiedToken, err := s.authClient.VerifyIDToken(ctx, token)
	if err != nil {
		return ctx, user.Unauthorized("invalid token")
	}

	log.Printf("Verified ID token: %v\n", verifiedToken)

	return ctx, nil
}
