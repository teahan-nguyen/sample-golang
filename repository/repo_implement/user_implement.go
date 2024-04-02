package repo_implement

import (
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
func (u *UserImplement) GetAllUsers(context context.Context) ([]*model.User, error) {
	collection := u.mongoDB.Collection("user")

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

func (u *UserImplement) GetUserById(context context.Context, userId string) (*model.User, error) {
	collection := u.mongoDB.Collection("user")
	var user *model.User
	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	if err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&user); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserImplement) UpdateUserById(context context.Context, userId string, input request.UpdateUser) (*model.User, error) {
	collection := u.mongoDB.Collection("user")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return nil, err
	}

	update := bson.M{
		"$set": bson.M{
			"email": input.Email,
			"role":  input.Role,
		},
	}

	_, err = collection.UpdateOne(context, bson.M{"_id": objectId}, update)
	if err != nil {
		return nil, err
	}

	var user *model.User
	err = collection.FindOne(context, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserImplement) RemoveUserById(context context.Context, userId string) error {
	collection := u.mongoDB.Collection("user")

	objectId, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(context, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	return nil
}
