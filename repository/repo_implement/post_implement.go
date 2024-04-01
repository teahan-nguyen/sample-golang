package repo_implement

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/repository"
)

type PostImplement struct {
	mongoDB *mongo.Database
}

func NewPostImplement(mongoDb *mongo.Database) repository.IPostRepository {
	return &PostImplement{
		mongoDB: mongoDb,
	}
}

func (a *PostImplement) CreatedPost(context context.Context, data request.ReqPost, userId string) (*response.CommonPostResponse, error) {
	collection := a.mongoDB.Collection("content")

	reqPost := &response.CommonPostResponse{
		ID:          primitive.NewObjectID(),
		Title:       data.Title,
		Description: data.Description,
		UserId:      userId,
	}

	insertData, err := collection.InsertOne(context, reqPost)
	if err != nil {
		return nil, err
	}

	objectID, ok := insertData.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	res := &response.CommonPostResponse{
		ID:          objectID,
		Title:       data.Title,
		Description: data.Description,
		UserId:      userId,
	}

	return res, nil
}

func (a *PostImplement) GetAllPosts(context context.Context) ([]*response.CommonPostResponse, error) {
	collection := a.mongoDB.Collection("content")

	cursor, err := collection.Find(context, bson.D{})
	if err != nil {
		return nil, err
	}

	var posts []*response.CommonPostResponse
	for cursor.Next(context) {
		var post response.CommonPostResponse
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (a *PostImplement) GetPostById(context context.Context, postId string) (*response.CommonPostResponse, error) {
	collection := a.mongoDB.Collection("content")

	var post *response.CommonPostResponse
	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return post, nil
}
func (a *PostImplement) RemovePostById(context context.Context, postId string, userId string) error {
	collection := a.mongoDB.Collection("content")

	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}

	var post struct {
		ID     primitive.ObjectID `bson:"_id"`
		UserId string             `bson:"userId"`
	}

	if err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&post); err != nil {
		return err
	}

	if post.UserId != userId {
		return errors.New("You don't have permission to delete this post")
	}
	_, err = collection.DeleteOne(context, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

func (a *PostImplement) UpdatePostById(context context.Context, postId string, input request.ReqUpdatePost, userId string) (*response.CommonPostResponse, error) {
	collection := a.mongoDB.Collection("content")

	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	var post struct {
		ID     primitive.ObjectID `bson:"_id"`
		UserId string             `bson:"userId"`
	}

	if err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&post); err != nil {
		return nil, err
	}

	if post.UserId != userId {
		return nil, errors.New("you don't have permission to delete this post")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       input.Title,
			"description": input.Description,
		},
	}

	_, err = collection.UpdateOne(context, bson.M{"_id": objectId}, update)
	if err != nil {
		return nil, err
	}

	var updatedPost *response.CommonPostResponse
	if err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&updatedPost); err != nil {
		return nil, err
	}

	return updatedPost, nil
}
