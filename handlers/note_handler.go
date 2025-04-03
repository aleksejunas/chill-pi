// This file defines a Go HTTP handler for managing notes in a web application. Here's a breakdown:
//
// 1. **Package and Imports**:
//    - The file is part of the `handlers` package.
//    - It imports necessary packages for JSON encoding, HTTP handling, and custom utilities/models.
//
// 2. **Fake Database**:
//    - A hardcoded slice of `models.Note` is used as a mock database to store notes.
//
// 3. **Handler Function**:
//    - `GetNotes`: Handles HTTP GET requests to the `/notes` endpoint.
//      - Logs the request using a utility function.
//      - Sets the response content type to JSON.
//      - Encodes and sends the `notes` slice as a JSON response.
//
// This is a simple implementation for serving a list of notes via an HTTP API.

// package handlers
//
// import (
//
//	"encoding/json"
//	"net/http"
//
//	"chill-pi/models"
//	"chill-pi/utils"
//
// )
//
// // Hardcoded slice of notes (fake database)
//
//	var notes = []models.Note{
//		{ID: 1, Title: "First Note", Content: "Greetings, Planet"},
//		{ID: 2, Title: "Second Note", Content: "Greetings, Universe"},
//	}
//
// // GET /notes
//
//	func GetNotes(w http.ResponseWriter, r *http.Request) {
//		utils.Info("GET /notes/")
//
//		w.Header().Set("Content-Type", "application/json")
//		json.NewEncoder(w).Encode(notes)
//	}
package handlers

import (
	"fmt"
	"net/http"

	"chill-pi/models"
	"chill-pi/utils"

	"github.com/gin-gonic/gin"
)

var notes = []models.Note{
	{ID: 1, Title: "First Note", Content: "Greetings, Planet"},
	{ID: 2, Title: "Second Note", Content: "Greetings, Universe"},
}

func GetNotes(c *gin.Context) {
	utils.Info("GET /notes")
	c.JSON(http.StatusOK, notes)
}

func GetNoteById(c *gin.Context) {
	idParam := c.Param("id")

	for _, note := range notes {
		if fmt.Sprintf("%d", note.ID) == idParam {
			utils.Info("GET /notes/" + idParam)
			c.JSON(http.StatusOK, note)
			return
		}
	}
}
