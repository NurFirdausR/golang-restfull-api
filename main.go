package main

import (
	"net/http"

	"github.com/NurFirdausR/golang-restfull-api/app"
	"github.com/NurFirdausR/golang-restfull-api/controller"
	"github.com/NurFirdausR/golang-restfull-api/helper"
	"github.com/NurFirdausR/golang-restfull-api/repository"
	"github.com/NurFirdausR/golang-restfull-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {

	validate := validator.New()
	db := app.NewDb()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories/", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err := server.ListenAndServe()

	helper.PanicIfErr(err)

}
