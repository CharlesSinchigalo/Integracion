package models

type Estudiante struct {
	CI      string `gorm:"primaryKey"` // cédula del estudiante
	Nombres string
}
