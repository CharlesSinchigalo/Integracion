package models

type Profesor struct {
	CI      string `gorm:"primaryKey"` // cédula del profesor
	Nombres string
}
