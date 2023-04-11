/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Mon Apr 10 2023 11:25:14 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package repository

import (
	"github.com/eulbyvan/app-mahasiswa/entity"
	"github.com/jmoiron/sqlx"
)

type StudentRepo interface {
	GetAll() []entity.Student
	GetById(id int) entity.Student
	Create(newStudent entity.Student) string
	UpdateById(id int, newStudent entity.Student) string
	DeleteById(id int) string
}

type studentRepo struct {
	// db []entity.Student
	db *sqlx.DB
}

func (s *studentRepo) GetAll() []entity.Student {
	// return s.db
	var students []entity.Student

	query := "SELECT id, name, age, major FROM students"
	rows, err := s.db.Query(query)
	if err != nil {
		return students
	}

	defer rows.Close()

	for rows.Next() {
		// fmt.Println("haloooo")
		var student entity.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Age, &student.Major)

		if err != nil {
			return students
		}

		students = append(students, student)
	}

	return students
}

func (s *studentRepo) GetById(id int) entity.Student {
	// for _, student := range s.db {
	// 	if student.ID == id {
	// 		return student
	// 	}
	// }

	return entity.Student{}
}

func (s *studentRepo) Create(newStudent entity.Student) string {
	// s.db = append(s.db, newStudent)
	return "create success"
}

func (s *studentRepo) UpdateById(id int, newStudent entity.Student) string {
	// for i, student := range s.db {
	// 	if student.ID == id {
	// 		s.db[i] = newStudent
	// 		return "update success"
	// 	}
	// }

	return "update failed"
}

func (s *studentRepo) DeleteById(id int) string {
	// for i, student := range s.db {
	// 	if student.ID == id {
	// 		s.db = append(s.db[:i], s.db[i+1:]...) // trik untuk hapus data di slice
	// 		return "delete success"
	// 	}
	// }

	return "delete failed"
}

func NewStudentRepo(db *sqlx.DB) StudentRepo {
	repo := new(studentRepo)
	repo.db = db
	return repo
}
