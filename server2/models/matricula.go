package models

type Matricula struct {
	ID           uint   `gorm:"primaryKey"`
	EstudianteCI string // FK a Estudiante
	AsignaturaID uint   // FK a Asignatura
	Ciclo        string
	N1           int
	N2           int
	Sup          int
}
