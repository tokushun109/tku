package main

import (
	"api/app/controllers"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StrStdin() (stringInput string) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	stringInput = scanner.Text()

	stringInput = strings.TrimSpace(stringInput)
	return
}

func main() {
	fmt.Println("販売サイトのurlを入力してください")
	url := StrStdin()
	controllers.DuplicateProduct(url)
}
