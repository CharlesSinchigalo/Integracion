package routes

import (
	"github.com/gin-gonic/gin"
	"server2/database"
	"server2/models"
	"net/http"
)

func RegistrarMatricula(c *gin.Context) {
	var matricula models.Matricula
	if err := c.ShouldBindJSON(&matricula); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	database.DB.Create(&matricula)
	c.JSON(http.StatusOK, gin.H{"mensaje": "Matrícula registrada"})
}
