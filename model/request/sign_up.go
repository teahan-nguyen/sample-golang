package request

type SignUp struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
