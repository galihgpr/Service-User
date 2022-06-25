package view

type RespondUser struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
	Phone string `json:"phone" form:"phone"`
	Role  string `json:"role" form:"role"`
}

type RespondLogin struct {
	Token string `json:"token"`
	User  RespondUser
}
