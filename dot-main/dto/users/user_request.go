package usersdto

type CreateUserRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	Npp      string `json:"npp" form:"phone" validate:"required"`
	NppSup   string `json:"nppsup" form:"addres" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Npp      string `json:"npp" form:"phone" `
	NppSup   string `json:"nppsup" form:"addres"`
}
