package contact_request

type ContactUpdateRequest struct {
	Id     string `param:"id" validate:"required"`
	Name   string `json:"name" validate:"required"`
	Email  string `json:"email" validate:"required,email"`
	Phone  string `json:"phone" validate:"required,phoneValidation"`
	Gender string `json:"gender" validate:"required,genderValidation"`
}
