/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Mon Apr 10 2023 11:25:20 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package usecase

import (
	"github.com/eulbyvan/app-mahasiswa/entity"
	"github.com/eulbyvan/app-mahasiswa/repository"
	"github.com/eulbyvan/app-mahasiswa/validation"
)

type StudentUsecase interface {
	FindAll() []entity.Student
	FindById(id int) entity.Student
	Register(newStudent entity.Student) string
	Edit(id int, newStudent entity.Student) string
	Unreg(id int) string
}

type studentUsecase struct {
	studentRepo repository.StudentRepo
}

func (u *studentUsecase) FindAll() []entity.Student {
	// business logic bisa ditambahkan disini
	return u.studentRepo.GetAll()
}

func (u *studentUsecase) FindById(id int) entity.Student {
	// business logic bisa ditambahkan disini
	return u.studentRepo.GetById(id)
}

func (u *studentUsecase) Register(newStudent entity.Student) string {
	// business logic, contoh: umur minimal 17 tahun
	err := validation.ValidateStudent(newStudent)

	if len(u.studentRepo.GetAll()) >= 10 {
		return "maximum number of students reached"
	}

	if err != nil {
		return err.Error()
	}
	
	return u.studentRepo.Create(newStudent)
}

func (u *studentUsecase) Edit(id int, newStudent entity.Student) string {
	// business logic bisa ditambahkan disini
	return u.studentRepo.UpdateById(id, newStudent)
}

func (u *studentUsecase) Unreg(id int) string {
	// business logic bisa ditambahkan disini
	return u.studentRepo.DeleteById(id)
}

func NewStudentUsecase(studentRepo repository.StudentRepo) StudentUsecase {
	return &studentUsecase {
		studentRepo: studentRepo,
	}
}