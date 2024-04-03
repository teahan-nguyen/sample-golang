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

func (p *PostImplement) CreatedPost(context context.Context, data request.RequestPost, userId string) (*response.CommonPostResponse, error) {
	collection := p.mongoDB.Collection("content")

	requestPost := &response.CommonPostResponse{
		ID:          primitive.NewObjectID(),
		Title:       data.Title,
		Description: data.Description,
		UserId:      userId,
	}

	insertData, err := collection.InsertOne(context, requestPost)
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

func (p *PostImplement) GetAllPosts(context context.Context) ([]*response.CommonPostResponse, error) {
	collection := p.mongoDB.Collection("content")

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

func (p *PostImplement) GetPostById(context context.Context, postId string) (*response.CommonPostResponse, error) {
	collection := p.mongoDB.Collection("content")

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
func (p *PostImplement) RemovePostById(context context.Context, postId string, userId string) error {
	if err := p.CheckPostPermission(context, postId, userId); err != nil {
		return err
	}

	collection := p.mongoDB.Collection("content")

	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}

func (p *PostImplement) UpdatePostById(context context.Context, postId string, input request.UpdatePost, userId string) (*response.CommonPostResponse, error) {
	if err := p.CheckPostPermission(context, postId, userId); err != nil {
		return nil, err
	}

	collection := p.mongoDB.Collection("content")

	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
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

func (p *PostImplement) CheckPostPermission(context context.Context, postId string, userId string) error {
	collection := p.mongoDB.Collection("content")

	objectId, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}

	var post struct {
		UserID string `bson:"userId"`
	}

	if err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&post); err != nil {
		return err
	}

	if post.UserID != userId {
		return errors.New("you don't have permission to delete this post")
	}

	return nil
}
