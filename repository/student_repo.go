package repository

import "github.com/eulbyvan/app-mahasiswa/entity"

type StudentRepo interface {
	GetAll() []entity.Student
	GetById(id int) entity.Student
	Create(newStudent entity.Student) string
	UpdateById(id int, newStudent entity.Student) string
	DeleteById(id int) string
}

type studentRepo struct {
	db []entity.Student
}

func (s *studentRepo) GetAll() []entity.Student {
	return s.db
}

func (s *studentRepo) GetById(id int) entity.Student {
	for _, student := range s.db {
		if student.ID == id {
			return student
		}
	}

	return entity.Student{}
}

func (s *studentRepo) Create(newStudent entity.Student) string {
	s.db = append(s.db, newStudent)
	return "create success"
}

func (s *studentRepo) UpdateById(id int, newStudent entity.Student) string {
	for i, student := range s.db {
		if student.ID == id {
			s.db[i] = newStudent
			return "update success"
		}
	}

	return "update failed"
}

func (s *studentRepo) DeleteById(id int) string {
	for i, student := range s.db {
		if student.ID == id {
			s.db = append(s.db[:i], s.db[i+1:]...) // trik untuk hapus data di slice
			return "delete success"
		}
	}

	return "delete failed"
}

func NewStudentRepo() StudentRepo {
	repo := new(studentRepo)
	student01 := entity.Student{
		ID: 1,
		Name: "Ismail",
		Age: 24,
		Major: "Computer Science",
	}
	student02 := entity.Student{
		ID: 2,
		Name: "Yoga",
		Age: 23,
		Major: "Information Technology",
	}
	student03 := entity.Student{
		ID: 3,
		Name: "Russel",
		Age: 26,
		Major: "Mechanical Engineering",
	}

	repo.db = []entity.Student{student01, student02, student03}
	return repo
}
