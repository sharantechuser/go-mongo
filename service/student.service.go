package service

import "github.com/sharantechuser/gomongo/model"

type IStudentService interface {
	GetAllStudents() ([]model.Student, error)

	CreateStudent(student model.Student) error
	UpdateStudent(_id string, student model.Student) (int64, error)
	DeleteStudent(id string, i interface{}) (int64, error)
}
