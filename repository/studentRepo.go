package repository

import (
	"svc-mahasiswa/models"
)

type Student interface {
	GetStudent(int)([]models.Student, error)
	GetStudents()([]models.Student, error)
	AddStudent(models.Student)(models.Student, error)
	UpdateStudent(int, models.Student)(models.Student, error)
	DeleteStudent(int) error
}