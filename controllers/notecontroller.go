package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hculpan/rpgnotes/config"
	"github.com/hculpan/rpgnotes/models"
)

type noteParams struct {
	Title    string
	Note     string
	Category string
	Keywords string
}

func NoteCreate(c *gin.Context) {
	noteParam := noteParams{}

	c.Bind(&noteParam)

	cat, err := models.FindCategory(noteParam.Category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	n := models.Note{
		Title:      noteParam.Title,
		Note:       noteParam.Note,
		CategoryID: cat.ID,
		Category:   *cat,
		Keywords:   noteParam.Keywords,
	}
	result := config.DB.Create(&n)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"note": n,
	})
}

func NotesIndex(c *gin.Context) {
	var notes []models.Note
	result := config.DB.Find(&notes)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	for i := 0; i < len(notes); i++ {
		cat, _ := models.FindCategoryByID(notes[i].CategoryID)
		notes[i].Category = *cat
	}

	c.JSON(http.StatusOK, gin.H{
		"count": len(notes),
		"notes": notes,
	})
}

func NoteShow(c *gin.Context) {
	id := c.Param("id")

	var note models.Note
	result := config.DB.First(&note, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	cat, _ := models.FindCategoryByID(note.CategoryID)
	note.Category = *cat

	c.JSON(http.StatusOK, gin.H{
		"notes": note,
	})
}

func NoteUpdate(c *gin.Context) {
	id := c.Param("id")

	noteParam := noteParams{}
	c.Bind(&noteParam)

	var n models.Note
	result := config.DB.First(&n, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	cat, err := models.FindCategory(noteParam.Category)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&n).Updates(
		models.Note{
			Title:      noteParam.Title,
			Note:       noteParam.Note,
			CategoryID: cat.ID,
			Keywords:   noteParam.Keywords,
		})

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	n.Category = *cat

	c.JSON(http.StatusOK, gin.H{
		"note": n,
	})
}

func NoteDelete(c *gin.Context) {
	id := c.Param("id")

	result := config.DB.Delete(&models.Note{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
