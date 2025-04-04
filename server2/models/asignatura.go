package models

type Asignatura struct {
	IDAsignatura uint   `gorm:"primaryKey"` // ID de la asignatura
	Asignatura   string // nombre
	Nivel        string // nivel
}
