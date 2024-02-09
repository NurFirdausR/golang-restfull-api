package helper

import (
	"github.com/NurFirdausR/golang-restfull-api/model/domain"
	"github.com/NurFirdausR/golang-restfull-api/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}
