package input

// CreateUserInput
type CreateUserInput struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=12"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
}

// UpdateUserInput
type UpdateUserInput struct {
	ID       int64  `json:"id" form:"username" validate:"required"`
	Username string `json:"username" form:"username" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required,email"`
}

// UserListInput
type UserListInput struct {
	Page     int    `json:"page" form:"page" validate:"required"`
	PageSize int    `json:"page_size" form:"page_size" validate:"required"`
	Username string `json:"username" form:"username"`
}
