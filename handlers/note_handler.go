/*
The `note_handler.go` file is part of the `handlers` package in a Go web application. It defines handler functions for managing HTTP requests related to notes. Here's a breakdown of its functionality:

### Purpose:
This file provides HTTP handler functions, `GetNoteById` and `GetNotes`, which are responsible for retrieving notes from the database. It uses the Gin web framework for routing and request handling.

### Key Components:
1. **Imports**:
   - `net/http`: Provides HTTP status codes and utilities.
   - `chill-pi/database`: Manages database interactions.
   - `chill-pi/models`: Defines the `Note` struct, representing a note in the database.
   - `chill-pi/utils`: Provides utility functions like logging.
   - `github.com/gin-gonic/gin`: A web framework for handling HTTP requests.

2. **Handler Functions**:
   - `GetNotes(c *gin.Context)`:
     - Queries the database for all notes using `database.DB.Find`.
     - Currently, it retrieves the notes but does not yet return them in the response.
   - `GetNoteById(c *gin.Context)`:
     - Extracts the `id` parameter from the URL.
     - Queries the database for a note with the given ID using `database.DB.First`.
     - If the note is not found, it logs a warning and responds with a `404 Not Found` status and an error message.
     - If the note is found, it logs the request (though the response is not yet implemented).

### Workflow:
1. A client sends a GET request to an endpoint like `/notes` or `/notes/:id`.
2. The appropriate handler function is invoked:
   - `GetNotes` retrieves all notes from the database.
   - `GetNoteById` retrieves a specific note based on the provided ID.
3. Both functions currently log the request but do not yet return the retrieved data in the response.

### Missing Implementation:
- The functions do not currently return the retrieved notes in the response. This should be added to complete the functionality.
*
*/

package handlers

import (
	"net/http"

	"chill-pi/database"
	"chill-pi/models"
	"chill-pi/utils"

	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var input models.Note

	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Warn("Invalid input for POST /notes")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	if input.Title == "" {
		utils.Warn("Title empty in POST /package.json")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	result := database.DB.Create(&input)
	if result.Error != nil {
		utils.Error("Failed to crate note: " + result.Error.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the note"})
		return
	}

	utils.Info("Created note: " + input.Title)
	c.JSON(http.StatusCreated, input)

}

func GetNotes(c *gin.Context) {
	var notes []models.Note
	database.DB.Find(&notes)
}

func GetNoteById(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	result := database.DB.First(&note, id)

	if result.Error != nil {
		utils.Warn("note not found: " + id)
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	utils.Info("GET /notes/" + id)
}
