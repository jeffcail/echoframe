package input

// CreateUserInput
type CreateUserInput struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"password" form:"password" validate:"required,gte=6,lte=12"`
}
