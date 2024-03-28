package response

import "go.mongodb.org/mongo-driver/bson/primitive"

type ResPostData struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	Title       string             `json:"title" bson:"title"`
	Description string             `json:"description"bson:"description"`
}
type PostResponse struct {
	TotalPages int              `json:"totalPages"`
	TotalItems int              `json:"totalItems"`
	Docs       []*PostDataEntry `json:"docs"`
}

type PostDataEntry struct {
	ID    primitive.ObjectID `json:"id"`
	Title string             `json:"title"`
	Desc  string             `json:"description"`
}
