/*
This file, `routes/routes.go`, defines the routing logic for a Gin-based web server in the `chill-pi` project. It registers HTTP endpoints and associates them with specific handler functions.

### Key Details:
1. **Package and Imports**:
  - The file belongs to the `routes` package.
  - It imports the `handlers` package for handling business logic and the `gin-gonic/gin` package for routing.

2. **Functionality**:
  - The `RegisterRoutes` function takes a `*gin.Engine` instance as input and sets up the following routes:
  - **Root Route (`/`)**: Responds with a JSON message welcoming users to the API.
  - **Notes Routes (`/notes`)**:
  - `GET /notes`: Calls `handlers.GetNotes` to retrieve all notes.
  - `GET /notes/:id`: Calls `handlers.GetNoteById` to retrieve a specific note by its ID.

3. **Usage**:
  - This function is intended to be called during server initialization to configure the API's routing structure.
*/

package routes

import (
	"chill-pi/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Velkommen til chill-pi API 👋"})
	})

	RegisterResourceRoutes(r, "/notes", handlers.NotesResource{})
}
