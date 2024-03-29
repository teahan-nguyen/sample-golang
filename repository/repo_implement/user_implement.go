package repo_implement

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"samples-golang/model"
	"samples-golang/model/request"
	"samples-golang/repository"
)

type UserImplement struct {
	mongoDB *mongo.Database
}

func NewUserImplement(mongoDb *mongo.Database) repository.IUserRepository {
	return &UserImplement{
		mongoDB: mongoDb,
	}
}
func (a *UserImplement) GetAllUsers(context context.Context) ([]*model.User, error) {
	collection := a.mongoDB.Collection("user")

	cursor, err := collection.Find(context, bson.D{})
	if err != nil {
		return nil, err
	}

	var users []*model.User
	for cursor.Next(context) {

		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (a *UserImplement) GetUserById(context context.Context, userId string) (*model.User, error) {
	collection := a.mongoDB.Collection("user")
	var user *model.User
	userObject, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(context, bson.M{"_id": userObject}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (a *UserImplement) UpdateUserById(context context.Context, id string, input request.UpdateUser) (*model.User, error) {
	collection := a.mongoDB.Collection("user")

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"email": input.Email,
			"role":  input.Role,
		},
	}
	fmt.Println(input.Email, input.Role)

	_, err = collection.UpdateOne(context, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	var updatedUser *model.User
	err = collection.FindOne(context, bson.M{"_id": objID}).Decode(&updatedUser)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (a *UserImplement) RemoveRoomById(context context.Context, id string) error {
	collection := a.mongoDB.Collection("user")

	userObject, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context, bson.M{"_id": userObject})
	if err != nil {
		return err
	}

	return nil
}
