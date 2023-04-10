package main

import (
	"fmt"

	"github.com/eulbyvan/app-mahasiswa/entity"
	"github.com/eulbyvan/app-mahasiswa/repository"
	"github.com/eulbyvan/app-mahasiswa/usecase"
)

func main() {
	studentRepo := repository.NewStudentRepo()
	studentUsecase := usecase.NewStudentUsecase(studentRepo)

	fmt.Println("==> List Before CRUD <==")
	fmt.Println(studentUsecase.FindAll())

	// tambahkan data mahasiswa
	fmt.Println("==> Register <==")
	fmt.Println(studentUsecase.Register(entity.Student{4, "test", 17, "test"}))

	// tampilkan data mahasiswa
	fmt.Println("==> Find All <==")
	fmt.Println(studentUsecase.FindAll())

	// tampilkan berdasarkan ID
	fmt.Println("==> Find By Id<==")
	fmt.Println(studentUsecase.FindById(2))

	// update berdasarkan ID
	fmt.Println("==> Edit By Id <==")
	fmt.Println(studentUsecase.Edit(4, entity.Student{4, "test01", 18, "test01"}))

	// delete berdasarkan ID
	fmt.Println("==> Delete By Id <==")
	fmt.Println(studentUsecase.Unreg(3))

	fmt.Println("==> List After Delete <==")
	fmt.Println(studentUsecase.FindAll())
}