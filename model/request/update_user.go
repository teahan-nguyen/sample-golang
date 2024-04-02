package request

import "github.com/go-playground/validator/v10"

type UpdateUser struct {
	Email string `json:"email" bson:"email"`
	Role  string `json:"role" bson:"role"`
}

func (input *UpdateUser) Validate() error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	return nil
}
