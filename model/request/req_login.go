package request

type ReqLogin struct {
	Email    string `validate:"required"`
	Password string `validate:"required"`
}
