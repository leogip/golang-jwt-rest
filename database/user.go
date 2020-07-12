package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/leogip/golang-jwt-rest/models"
)

// FindAllUser ...
func FindAllUser() ([]models.User, error) {
	var users []models.User
	cursor, err := CollectionUser.Find(context.TODO(), bson.D{{}})
    if err !=nil {
        log.Fatal(err)
	}
	
	for cursor.Next(context.TODO()) {
        var elem models.User
		err := cursor.Decode(&elem)
		elem.Password = ""
        if err != nil {
            log.Fatal(err)
        }

        users = append(users, elem)

    }

    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }

    cursor.Close(context.TODO())

	return users, err
}

// FindUserByID ...
func FindUserByID(id primitive.ObjectID) (models.User, error) {
	var user models.User
	err := CollectionUser.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

// InsertUser ...
func InsertUser(user models.User) error {
	_, err := CollectionUser.InsertOne(context.TODO(), &user)
	return err
}

// FindByEmail ...
func FindByEmail(email string) (models.User, error) {
	var user models.User
	
	err := CollectionUser.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	return user, err
}