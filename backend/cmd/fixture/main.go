package main

import (
	"context"
	"fmt"
	"golang-nextjs-app/tools"
	"log"
	"os"
)

const (
	auth = "create_auth_user"
	task = "create_task"
)

func main() {
	args := os.Args
	ctx := context.Background()

	if len(args) < 2 {
		fmt.Println("引数が指定されていません")
		return
	}

	arg := args[1]

	switch arg {
	case auth:
		log.Println("auth user を作成します")
		tools.CreateDummyAuthUser(ctx)
	case task:
		log.Println("task を作成します")
		tools.CreateDummyTask(ctx)
	default:
		fmt.Println("不明な引数です")
	}

}
