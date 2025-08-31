package main

import (
	"api/app/controllers"
	_ "api/config"
)

func main() {
	controllers.StartMainServer()
}
