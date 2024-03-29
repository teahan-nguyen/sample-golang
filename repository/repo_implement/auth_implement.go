package repo_implement

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"samples-golang/model"
	"samples-golang/repository"
)

type AuthRepository struct {
	mongoDB *mongo.Database
}

func NewAuthRepository(mongoDb *mongo.Database) repository.IAuthRepository {
	return &AuthRepository{
		mongoDB: mongoDb,
	}
}

func (a AuthRepository) InsertUser(context context.Context, email string) (*model.User, error) {
	collection := a.mongoDB.Collection("user")

	filter := bson.M{"email": email}
	var existingUser model.User
	err := collection.FindOne(context, filter).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		count, err := collection.CountDocuments(context, bson.M{})
		if err != nil {
			return nil, err
		}

		newUser := model.User{
			Email: email,
		}
		newUser.SetRoleByCount(int(count))
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

func (a AuthRepository) VerifyUser(context context.Context, email string) (*model.User, error) {
	var user model.User
	collection := a.mongoDB.Collection("user")

	filter := bson.M{"email": email}
	err := collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &model.User{}, errors.New("user not found")
		}
		return &model.User{}, err
	}

	return &user, nil
}
