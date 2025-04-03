package main

import (
	"log"
	"net/http"

	"chill-pi/utils"
)

func main() {
	utils.InitLogger("logs/server.log")
	defer utils.CloseLogger()

	utils.Info("Starter server på http://localhost:8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Serveren kjører 👋"))
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start the server: %v", err)
	}
}
