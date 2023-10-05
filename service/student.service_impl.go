package service

import (
	"github.com/sharantechuser/gomongo/db"
	"github.com/sharantechuser/gomongo/model"
)

const (
	student_coll_name = "Student"
)

type StudentServiceImpl struct {
}

func NewStudentService() IStudentService {
	return &StudentServiceImpl{}
}

func (us *StudentServiceImpl) GetAllStudents() ([]model.Student, error) {

	var results []model.Student
	result, err := db.Handle.GetAll(student_coll_name, results)

	if err != nil {
		return []model.Student{}, err
	}

	return result.([]model.Student), err
}

func (us *StudentServiceImpl) CreateStudent(student model.Student) error {
	return db.Handle.Create(student)
}
func (us *StudentServiceImpl) UpdateStudent(_id string, student model.Student) (int64, error) {
	return db.Handle.Update(_id, student)
}
func (us *StudentServiceImpl) DeleteStudent(_id string, student interface{}) (int64, error) {
	return db.Handle.Update(_id, student)
}
