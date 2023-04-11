/*
 * Author : Ismail Ash Shidiq (https://eulbyvan.netlify.app)
 * Created on : Mon Apr 10 2023 11:25:29 AM
 * Copyright : Ismail Ash Shidiq © 2023. All rights reserved
 */

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/eulbyvan/app-mahasiswa/repository"
	"github.com/eulbyvan/app-mahasiswa/usecase"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	// db connection
	dbHost := "localhost"
	dbPort := "5432"
	dbUser := "postgres"
	dbPassword := "postgres"
	dbName := "testdb"
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	db, err := sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected")
	}
	
	defer func() {
		err := db.Close()
		if err != nil {
			panic(err)
		}
		log.Println("Connection to db is closed")
	}()

	studentRepo := repository.NewStudentRepo(db)
	studentUsecase := usecase.NewStudentUsecase(studentRepo)

	os.Setenv("MIN_STUDENT_AGE", "17")

	fmt.Println("==> List Before CRUD <==")
	fmt.Println(studentUsecase.FindAll())

	// // tambahkan data mahasiswa
	// fmt.Println("==> Register <==")
	// fmt.Println(studentUsecase.Register(entity.Student{4, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{5, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{6, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{7, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{8, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{9, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{10, "test", 17, "test"}))
	// fmt.Println(studentUsecase.Register(entity.Student{11, "test", 17, "test"}))

	// // tampilkan data mahasiswa
	// fmt.Println("==> Find All <==")
	// fmt.Println(studentUsecase.FindAll())

	// // tampilkan berdasarkan ID
	// fmt.Println("==> Find By Id<==")
	// fmt.Println(studentUsecase.FindById(2))

	// // update berdasarkan ID
	// fmt.Println("==> Edit By Id <==")
	// fmt.Println(studentUsecase.Edit(4, entity.Student{4, "test01", 18, "test01"}))

	// // delete berdasarkan ID
	// fmt.Println("==> Delete By Id <==")
	// fmt.Println(studentUsecase.Unreg(3))

	// fmt.Println("==> List After Delete <==")
	// fmt.Println(studentUsecase.FindAll())
}