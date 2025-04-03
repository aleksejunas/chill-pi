package main

import (
	"net/http"

	// "chill-pi/handlers"
	"chill-pi/routes"
	"chill-pi/utils"
)

func main() {
	utils.InitLogger("logs/server.log")
	defer utils.CloseLogger()

	utils.Info("Starter server på http://localhost:8080")

	utils.Info("Starting chill-pi API on http://localhost:8080")

	routes.RegisterRoutes()

	// http.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
	// 	if r.Method == http.MethodGet {
	// 		handlers.GetNotes(w, r)
	// 	} else {
	// 		http.Error(w, "Metode ikke støttet", http.StatusMethodNotAllowed)
	// 	}
	// })

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		utils.Error("Server Error: " + err.Error())
	}
}
