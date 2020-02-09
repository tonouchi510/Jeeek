package service_test

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"fmt"
	"google.golang.org/api/option"
	"log"
	"os"
	"testing"
)

var fsClient *firestore.Client
var ctx context.Context
var FirebaseCredentials = os.Getenv("FIREBASE_CREDENTIALS")

func TestMain(m *testing.M) {
	var exitStatus int
	defer func() {
		os.Exit(exitStatus)
	}()

	projectID := "jeeek-dev"
	ctx = context.Background()

	var err error
	_, _ = InitFirebaseAuth(ctx)
	fsClient, err = firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("Failed to create firestore client: %v", err)
	}
	defer fsClient.Close()

	fmt.Println("Test start...")
	exitStatus = m.Run()
}

func InitFirebaseAuth(ctx context.Context) (app *firebase.App, client *auth.Client) {
	opt := option.WithCredentialsFile(FirebaseCredentials)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	// Get an auth client from the firebase.App
	client, err = app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return app, client
}
