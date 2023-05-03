package firebase

import (
	"context"
	"fmt"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

var firebaseApp *firebase.App

func InitFirebase() {
	opt := option.WithCredentialsFile("backend/config/golang-nextjs-app-firebase-adminsdk-xbyxa-6027c29d3d.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	firebaseApp = app
}

func VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	client, err := firebaseApp.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v", err)
	}
	return client.VerifyIDToken(ctx, idToken)
}
