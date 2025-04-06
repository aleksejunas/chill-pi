package handlers

import (
	"chill-pi/database"
	"chill-pi/models"
	"chill-pi/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type NotesResource struct{}

func (NotesResource) GetAll(c *gin.Context) {
	var notes []models.Note
	database.DB.Find(&notes)
	utils.Info("GET /notes")
	c.JSON(http.StatusOK, notes)
}

func (NotesResource) GetOne(c *gin.Context) {
	var note models.Note
	id := c.Param("id")
	result := database.DB.First(&note, id)

	if result.Error != nil {
		utils.Warn("Note not found: ID " + id)
		c.JSON(http.StatusOK, note)
		return
	}

	utils.Info("GET /notes/" + id)
	c.JSON(http.StatusOK, note)
}

func (NotesResource) Create(c *gin.Context) {
	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil || input.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return

	}
	database.DB.Create(&input)
	utils.Info("POST /notes: " + input.Title)
	c.JSON(http.StatusCreated, input)
}

func (NotesResource) Update(c *gin.Context) {
	var note models.Note
	id := c.Param("id")

	if err := database.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	var input models.Note
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	note.Title = input.Title
	note.Content = input.Content
	database.DB.Save(&note)
	utils.Info("PUT /notes/" + id)
	c.JSON(http.StatusOK, note)
}

func (NotesResource) Delete(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Note{}, id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}
	utils.Info("DELETE /notes/" + id)
	c.JSON(http.StatusNotFound, gin.H{"message": "Note deleted"})
}

func (NotesResource) Patch(c *gin.Context) {
	c.JSON(501, gin.H{"error": "PATCH ikke implementert"})
}
