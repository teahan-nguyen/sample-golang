package repo_implement

import (
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"samples-golang/db"
	"samples-golang/model"
	"samples-golang/model/request"
	"samples-golang/model/response"
	"samples-golang/repository"
)

type AuthImplement struct {
	mongDB *db.MongoDB
}

func NewImplement(mongDB *db.MongoDB) repository.AutheRepository {
	return &AuthImplement{
		mongDB: mongDB,
	}
}

func (a *AuthImplement) CreatedPost(context context.Context, data request.ReqPost, userId string) (*response.ResPostData, error) {
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("content")
	dataWithUserId := struct {
		request.ReqPost
		UserId string `bson:"userId"`
	}{
		ReqPost: data,
		UserId:  userId,
	}
	insertData, err := collection.InsertOne(context, dataWithUserId)
	if err != nil {
		return nil, err
	}
	objectID, ok := insertData.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.New("failed to convert InsertedID to ObjectID")
	}

	res := &response.ResPostData{
		ID:          objectID,
		Title:       data.Title,
		Description: data.Description,
	}

	return res, nil
}

func (a *AuthImplement) GetAllPosts(context context.Context) ([]*response.ResPostData, error) {
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("content")

	cursor, err := collection.Find(context, bson.D{})
	if err != nil {
		return nil, err
	}

	var posts []*response.ResPostData
	for cursor.Next(context) {
		var post bson.M
		if err := cursor.Decode(&post); err != nil {
			return nil, err
		}

		reqPost := post["reqpost"].(bson.M)
		title := reqPost["title"].(string)
		description := reqPost["description"].(string)
		id := post["_id"].(primitive.ObjectID).Hex()
		ObjectId, _ := primitive.ObjectIDFromHex(id)

		resPost := &response.ResPostData{
			ID:          ObjectId,
			Title:       title,
			Description: description,
		}

		posts = append(posts, resPost)
	}

	return posts, nil
}

func (a *AuthImplement) GetPostById(context context.Context, postId string) (*response.ResPostData, error) {
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("content")

	var post bson.M
	postObject, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context, bson.M{"_id": postObject}).Decode(&post)
	if err != nil {
		return nil, err
	}

	reqPost := post["reqpost"].(bson.M)
	title := reqPost["title"].(string)
	description := reqPost["description"].(string)

	resPost := &response.ResPostData{
		ID:          postObject,
		Title:       title,
		Description: description,
	}

	return resPost, nil
}
func (a *AuthImplement) RemovePostById(context context.Context, postId string, userId string) error {
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("content")

	postObject, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return err
	}

	var post struct {
		ID     primitive.ObjectID `bson:"_id"`
		UserId string             `bson:"userId"`
	}

	err = collection.FindOne(context, bson.M{"_id": postObject}).Decode(&post)
	if err != nil {
		return err
	}

	if post.UserId != userId {
		return errors.New("You don't have permission to delete this post")
	}
	_, err = collection.DeleteOne(context, bson.M{"_id": postObject})
	if err != nil {
		return err
	}

	return nil
}

func (a *AuthImplement) UpdatePostById(context context.Context, postId string, input response.ResPostData, userId string) (*response.ResPostData, error) {
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("content")

	postObject, err := primitive.ObjectIDFromHex(postId)
	if err != nil {
		return nil, err
	}

	var post struct {
		ID     primitive.ObjectID `bson:"_id"`
		UserId string             `bson:"userId"`
	}
	err = collection.FindOne(context, bson.M{"_id": postObject}).Decode(&post)
	if err != nil {
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

	_, err = collection.UpdateOne(context, bson.M{"_id": postObject}, update)
	if err != nil {
		return nil, err
	}
	var updatedPost *response.ResPostData
	err = collection.FindOne(context, bson.M{"_id": postObject}).Decode(&updatedPost)
	if err != nil {
		return nil, err
	}

	return updatedPost, nil
}

func (a *AuthImplement) InsertUser(context context.Context, email string) (*model.User, error) {
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("user")

	filter := bson.M{"email": email}
	var existingUser model.User
	err := collection.FindOne(context, filter).Decode(&existingUser)
	if err == mongo.ErrNoDocuments {
		count, err := collection.CountDocuments(context, bson.M{})
		if err != nil {
			return nil, err
		}
		var role string
		if count == 0 {
			role = "ADMIN"
		} else {
			role = "USER"
		}

		newUser := model.User{
			Email: email,
			Role:  role,
		}
		_, err = collection.InsertOne(context, newUser)
		if err != nil {
			return nil, err
		}
		return &newUser, nil
	} else if err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (a *AuthImplement) VerifyUser(context context.Context, email string) (*model.User, error) {
	var user model.User
	collection := a.mongDB.Client.Database(a.mongDB.DbName).Collection("user")

	filter := bson.M{"email": email}
	err := collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.User{}, fmt.Errorf("user not found")
		}
		return &model.User{}, err
	}

	return &user, nil
}
