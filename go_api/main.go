package main

import (
	"api/app/controllers"
	_ "api/app/models"
	_ "api/config"
)

func main() {
	controllers.StartMainServer()
}
