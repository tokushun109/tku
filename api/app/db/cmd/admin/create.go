package main

import (
	"api/app/models"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// go run app/db/cmd/admin/create.goで管理ユーザーを作成
func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	user := &models.User{}

	// 名前の入力
	fmt.Println("名前を入力してください")
	name := StrStdin()
	user.Name = name

	// メールアドレスの入力
	fmt.Println("メールアドレスを入力してください")
	email := StrStdin()
	user.Email = email

	// パスワードの入力
	fmt.Println("パスワードを入力してください")
	password := StrStdin()
	user.Password = password

	// userの作成用のメソッド
	models.InsertUser(user, true)

	fmt.Println("管理ユーザーを作成しました")
}
