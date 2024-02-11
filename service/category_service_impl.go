package service

import (
	"context"
	"database/sql"

	"github.com/NurFirdausR/golang-restfull-api/exception"
	"github.com/NurFirdausR/golang-restfull-api/helper"
	"github.com/NurFirdausR/golang-restfull-api/model/domain"
	"github.com/NurFirdausR/golang-restfull-api/model/web"
	"github.com/NurFirdausR/golang-restfull-api/repository"
	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, DB *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)

	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = service.CategoryRepository.Save(ctx, tx, category)
	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfErr(err)
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.CategoryRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundeError(err.Error()))
	}

	category.Id = request.Id
	category.Name = request.Name

	category = service.CategoryRepository.Update(ctx, tx, category)
	return helper.ToCategoryResponse(category)

}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundeError(err.Error()))
	}
	service.CategoryRepository.Delete(ctx, tx, category)
}

func (service *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)
	category, err := service.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundeError(err.Error()))
	}

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfErr(err)
	defer helper.CommitOrRollback(tx)
	categories := service.CategoryRepository.FindAll(ctx, tx)
	var CategoryResponses []web.CategoryResponse
	for _, category := range categories {
		CategoryResponses = append(CategoryResponses, helper.ToCategoryResponse(category))
	}
	return CategoryResponses
}
