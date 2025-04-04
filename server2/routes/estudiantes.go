package routes

import (
	"github.com/gin-gonic/gin"
	"server2/database"
	"server2/models"
	"net/http"
)

func RegistrarEstudiante(c *gin.Context) {
	var estudiante models.Estudiante
	if err := c.ShouldBindJSON(&estudiante); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&estudiante)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Estudiante registrado"})
}
