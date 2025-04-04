package main

import (
	"github.com/gin-gonic/gin"
	"server2/database"
	"server2/models"
	"server2/routes"
)

func main() {
	// Iniciar conexión a MySQL
	database.Connect()

	// Migrar todos los modelos a la base de datos
	database.DB.AutoMigrate(
		&models.Estudiante{},
		&models.Asignatura{},
		&models.Profesor{},
		&models.Matricula{},
		&models.ProfesorCiclo{},
	)

	// Crear router de Gin
	r := gin.Default()

	// Rutas del sistema
	r.POST("/estudiantes", routes.RegistrarEstudiante)
	r.POST("/asignaturas", routes.RegistrarAsignatura)
	r.POST("/profesores", routes.RegistrarProfesor)
	r.POST("/matriculas", routes.RegistrarMatricula)
	r.POST("/profesor-ciclo", routes.RegistrarProfesorCiclo)

	// Puerto de escucha del servidor
	r.Run(":8001")
}
