package database

import (
	"context"
	"demo/dto"
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SaveMessage(userID string, message string) {
	fmt.Println("MongoDB connecting...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dburi := viper.GetViper().GetString("database.URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Println("MongoDB connected.")

	database := client.Database("demo")
	collection := database.Collection("demo")
	var result dto.Messages
	err = collection.FindOne(context.TODO(), bson.M{"userID": userID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("No document, create new userID...")
			result.UserID = userID
			result.Message = append(result.Message, message)
			res, err := collection.InsertOne(context.Background(), result)
			if err != nil {
				panic(err)
			}
			fmt.Printf("res: %v\n", res)
			return
		}
	}

	newMessageArray := append(result.Message, message)
	update := bson.D{{"$set", bson.D{{"message", newMessageArray}}}}

	fmt.Printf("result: %v\n", result)
	updateResult, updateErr := collection.UpdateOne(context.TODO(),
		bson.D{{Key: "userID", Value: userID}},
		update)
	if err != nil {
		panic(updateErr)
	}
	fmt.Printf("updateResult: %v\n", updateResult)
}

func GetMessage(userID string) []string {
	fmt.Println("MongoDB connecting...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dburi := viper.GetViper().GetString("database.URI")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburi))

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	fmt.Println("MongoDB connected.")

	database := client.Database("demo")
	collection := database.Collection("demo")
	var result dto.Messages
	err = collection.FindOne(context.TODO(), bson.M{"userID": userID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Println("user not exist")
			return nil
		}
	}
	return result.Message
}
