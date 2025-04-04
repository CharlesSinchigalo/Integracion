package routes

import (
	"github.com/gin-gonic/gin"
	"server2/database"
	"server2/models"
	"net/http"
)

func RegistrarProfesor(c *gin.Context) {
	var profesor models.Profesor
	if err := c.ShouldBindJSON(&profesor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&profesor)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Profesor registrado"})
}
