package routes

import (
	"github.com/gin-gonic/gin"
	"server2/database"
	"server2/models"
	"net/http"
)

func RegistrarAsignatura(c *gin.Context) {
	var asignatura models.Asignatura
	if err := c.ShouldBindJSON(&asignatura); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&asignatura)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Asignatura registrada"})
}
