package main

import (
	"context"
	"errors"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// api-service check user is available
func (app *application) authenticateMetaToken(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md.Get("Authorization")) != 1 {
		return status.Error(codes.PermissionDenied, "error exist auth header")
	}
	token := md.Get("Authorization")[0]
	if err := app.authenticateToken(token); err != nil {
		return status.Error(codes.PermissionDenied, err.Error())
	}
	return nil
}

func (app *application) authenticateToken(authorizationHeader string) error {
	headerParts := strings.Split(authorizationHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return errors.New("no authorization header received")
	}

	token := headerParts[1]
	if len(token) != 26 {
		return errors.New("authentication token wrong size")
	}

	// get the user from the tokens table
	_, err := app.DB.GetUserForToken(token)
	if err != nil {
		return errors.New("no matching user found")
	}
	return nil
}
