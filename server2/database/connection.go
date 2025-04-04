package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// Cambia estos valores según tu configuración de MySQL
	username := "root"
	password := "123456"
	host := "127.0.0.1"
	port := "3306"
	dbname := "midb"

	// Cadena de conexión
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8mb4&parseTime=True&loc=Local"

	// Conectar a MySQL
	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Error al conectar con MySQL: ", err)
	}

	log.Println("✅ Conexión a MySQL exitosa")
	DB = connection
// Migrate the schema
}
