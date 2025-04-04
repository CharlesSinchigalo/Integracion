package models

type ProfesorCiclo struct {
	ID           uint   `gorm:"primaryKey"`
	ProfesorCI   string // FK a Profesor
	AsignaturaID uint   // FK a Asignatura
	Ciclo        string
}
