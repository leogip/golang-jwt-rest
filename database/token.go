package database

import (
	"context"

	"github.com/leogip/golang-jwt-rest/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// FindToken ...
func FindToken(token string) error {
	r := bson.D{}
	err := CollectionTokens.FindOne(context.TODO(), bson.M{"token": token}).Decode(r)
	return err
}

// FindTokenByUser ...
func FindTokenByUser(userID primitive.ObjectID) (models.Token, error) {
	var token models.Token
	err := CollectionTokens.FindOne(context.TODO(), bson.M{"userId": userID}).Decode(&token)
	return token, err
}

// FindAllTokens ...
func FindAllTokens() ([]models.Token, error) {
	var tokens []models.Token
	cursor, err := CollectionTokens.Find(context.TODO(), bson.D{{}})
    if err != nil {
        log.Fatal(err)
	}
	
	for cursor.Next(context.TODO()) {
        var elem models.Token
		err := cursor.Decode(&elem)
        if err != nil {
            log.Fatal(err)
        }

        tokens = append(tokens, elem)

    }

    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }

    cursor.Close(context.TODO())

	return tokens, err
}

// InsertToken ...
func InsertToken(token models.Token) error {
	_, err := CollectionTokens.InsertOne(context.TODO(), &token)
	return err
}

// UpdateToken ...
func UpdateToken(id primitive.ObjectID, token string) error {
	_, err := CollectionTokens.UpdateOne(
		context.TODO(),
		bson.M{"_id": id},
		bson.D{
			{"$set", bson.D{{"token", token}}},
		},
	)
	return err
}

// RemoveTokenByID ...
func RemoveTokenByID(id primitive.ObjectID) error {
	_, err := CollectionTokens.DeleteOne(context.TODO(), bson.M{"_id": id})
	return err
}

// RemoveTokenByUserID ...
func RemoveTokenByUserID(userID primitive.ObjectID) error {
	_, err := CollectionTokens.DeleteMany(context.TODO(), bson.M{"userId": userID})
	return err
}