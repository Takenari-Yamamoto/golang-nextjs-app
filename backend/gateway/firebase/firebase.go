package firebase

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// TODO: firebase.NewAppは外から渡したいかも
type Firebase struct{}

func NewFirebase() *Firebase {
	return &Firebase{}
}

func getFirebaseCredentialsFromSecretManager(ctx context.Context, secretName string) ([]byte, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	// Secret Managerから認証情報を取得
	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: secretName,
	}

	result, err := client.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		log.Default().Println("Secret Managerからの認証情報の取得に失敗しました: ", err)
		return nil, err
	}
	log.Default().Println("Secret Managerからの認証情報の取得に成功しました: ", result)

	return result.Payload.Data, nil
}

func InitFirebase() *firebase.App {
	ctx := context.Background()

	projectID := os.Getenv("PROJECT_ID")
	secretName := fmt.Sprintf("projects/%s/secrets/firebase_credentials/versions/latest", projectID)
	credentials, err := getFirebaseCredentialsFromSecretManager(ctx, secretName)
	if err != nil {
		log.Fatalf("failed to fetch firebase credentials: %v", err)
	}
	err = ioutil.WriteFile("config/firebase_credentials.json", credentials, 0644)
	if err != nil {
		log.Fatalf("failed to write secret to file: %v", err)
	}
	opt := option.WithCredentialsFile("config/firebase_credentials.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	return app
}

// func InitFirebase() *firebase.App {
// 	opt := option.WithCredentialsFile("config/firebase_credentials.json")
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	fmt.Printf("app: %+v\n", app)
// 	if err != nil {
// 		log.Fatalf("firebase.NewApp: %v", err)
// 	}
// 	return app
// }

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

// authUserを全件取得
func (f *Firebase) ListUsers(ctx context.Context) ([]*auth.ExportedUserRecord, error) {
	var users []*auth.ExportedUserRecord
	app := InitFirebase()
	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	var nextPageToken string
	for {
		// ユーザー情報を取得
		it := client.Users(ctx, nextPageToken)
		for {
			user, err := it.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				log.Fatalf("error listing users: %v\n", err)
			}
			users = append(users, user)

		}
		// 次のページがある場合、次のページのTokenを取得
		nextPageToken = it.PageInfo().Token
		if nextPageToken == "" {
			break
		}
	}
	return users, nil
}
