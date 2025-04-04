package main

import (
	"chill-pi/database"
	"chill-pi/routes"
	"chill-pi/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitLogger("logs/server.log")
	defer utils.CloseLogger()

	database.ConnectDB()

	r := gin.Default()

	routes.RegisterRoutes(r)

	utils.Info("Starter ChillPi med Gin på http://localhost:8080")
	r.Run(":8080")
}

// err := http.ListenAndServe(":8080", nil)
// if err != nil {
// 	utils.Error("Server Error: " + err.Error())
// }
// }
