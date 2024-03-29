package response

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommonPostResponse struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description"bson:"description"`
	UserId      string             `json:"userId" bson:"userId"`
}

type PostDataEntry struct {
	ID     primitive.ObjectID `json:"id"`
	Title  string             `json:"title"`
	Desc   string             `json:"description"`
	UserId string             `json:"userId"`
}

func (input *CommonPostResponse) Validate() error {
	validate := validator.New()
	if err := validate.Struct(input); err != nil {
		return err
	}

	return nil
}
