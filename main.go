package main

import (
	"net/http"

	"chill-pi/routes"
	"chill-pi/utils"
)

func main() {
	utils.InitLogger("logs/server.log")
	defer utils.CloseLogger()

	utils.Info("Starting chill-pi API on http://localhost:8080")

	routes.RegisterRoutes()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		utils.Error("Server Error: " + err.Error())
	}
}
