package view

type AddUser struct {
	Name     string `json:"name" validate:"required" form:"name"`
	Email    string `json:"email" validate:"required" form:"email"`
	Password string `json:"password" validate:"required" form:"password"`
	Phone    string `json:"phone" validate:"required" form:"phone"`
	Role     string `json:"role" validate:"required" form:"role"`
}

type UpdateUser struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Phone    string `json:"phone" form:"phone"`
	Role     string `json:"role" form:"role"`
}

type Login struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
