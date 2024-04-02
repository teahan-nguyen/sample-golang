package request

type Login struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
