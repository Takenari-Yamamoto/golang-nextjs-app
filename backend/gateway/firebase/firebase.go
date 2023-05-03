package firebase

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

// TODO: firebase.NewAppは外から渡したいかも
type Firebase struct{}

func NewFirebase() *Firebase {
	return &Firebase{}
}

func InitFirebase() *firebase.App {
	// よくない気もするが。。
	opt := option.WithCredentialsFile("config/firebase_credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("firebase.NewApp: %v", err)
	}
	return app
}

func (f *Firebase) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {
	app := InitFirebase()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	token, err := client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, err
	}
	return &auth.Token{
		UID: token.UID,
	}, nil
}

// UIDを元にユーザー情報を取得
func (f *Firebase) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	app := InitFirebase()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	user, err := client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// auth user を作成
func (f *Firebase) CreateUser(ctx context.Context, email, password string) (*auth.UserRecord, error) {
	app := InitFirebase()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	params := (&auth.UserToCreate{}).
		Email(email).
		Password(password)
	user, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return user, nil
}
