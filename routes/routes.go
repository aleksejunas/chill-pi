// This file defines a Go package named `routes` that registers HTTP routes for a web server. Specifically:
//
// - It imports the `net/http` package for handling HTTP requests and responses, and a custom `handlers` package for processing specific routes.
// - The `RegisterRoutes` function sets up an HTTP route `/notes` using `http.HandleFunc`.
// - When a request is made to `/notes`:
//   - If the HTTP method is `GET`, it calls the `handlers.GetNotes` function to handle the request.
//   - For any other HTTP method, it responds with a `405 Method Not Allowed` error and a message "Metode ikke støttet" (Norwegian for "Method not supported").
//
// This file is likely part of a web application where `RegisterRoutes` is called to configure the server's routing logic.

package routes

import (
	"net/http"

	"chill-pi/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Velkommen til chill-pi API 👋"}`))
	})

	http.HandleFunc("/notes/", func(w http.ResponseWriter, r *http.Request) {
		// Da matcher det både /notes og /notes/ og /notes/1 osv.
		if r.Method == http.MethodGet {
			handlers.GetNotes(w, r)
		} else {
			http.Error(w, "Metode ikke støttet", http.StatusMethodNotAllowed)
		}
	})
}
