package routes

import (
	"github.com/gin-gonic/gin"
	"server2/database"
	"server2/models"
	"net/http"
)

func RegistrarProfesorCiclo(c *gin.Context) {
	var pc models.ProfesorCiclo
	if err := c.ShouldBindJSON(&pc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&pc)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Profesor-Ciclo registrado"})
}
