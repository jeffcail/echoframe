package input

// AuthLoginInput
type AuthLoginInput struct {
	Username string `json:"username" form:"username" validate:"required"`
	Password string `json:"Password" form:"Password" validate:"required"`
}
