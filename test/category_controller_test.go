package test

import (
	"database/sql"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/NurFirdausR/golang-restfull-api/app"
	"github.com/NurFirdausR/golang-restfull-api/controller"
	"github.com/NurFirdausR/golang-restfull-api/helper"
	"github.com/NurFirdausR/golang-restfull-api/middleware"
	"github.com/NurFirdausR/golang-restfull-api/repository"
	"github.com/NurFirdausR/golang-restfull-api/service"
	"github.com/go-playground/validator/v10"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:Bismillah@123@tcp(localhost:3306)/golang_api_test")

	helper.PanicIfErr(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name" : "gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories/", requestBody)
	request.Header.Add("Content-Type", "application/json") // Adjusted Content-Type header
	request.Header.Add("X-API-KEY", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, http.StatusOK, response.StatusCode) // Adjusted status code to http.StatusOK
}

func TestCreateCategoryFailed(t *testing.T) {

}
func TestUpdateCategorySuccess(t *testing.T) {

}
func TestUpdateCategoryFailed(t *testing.T) {

}
func TestDeleteCategorySuccess(t *testing.T) {

}

func TestDeleteCategoryFailed(t *testing.T) {

}

func TestGetCategorySuccess(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}
func TestAllCategorySuccess(t *testing.T) {

}

func TestAllCategoryFailed(t *testing.T) {

}
