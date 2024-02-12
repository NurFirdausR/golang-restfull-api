package web

type CategoryCreateRequest struct {
	Name string `validate:"required,max=30,min=1" json:"name"`
}
