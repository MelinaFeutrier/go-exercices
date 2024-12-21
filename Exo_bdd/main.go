package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://MelinaFeu:123456..@cluster0.grdqi.mongodb.net/"))
	if err != nil {
		log.Fatal("Erreur de connexion :", err)
	}
	defer client.Disconnect(ctx)

	collection := client.Database("testdb").Collection("users")

	newUser := bson.D{
		{Key: "name", Value: "Melina Feutrier"},
		{Key: "email", Value: "melina@example.com"},
		{Key: "age", Value: 33},
	}
	insertResult, err := collection.InsertOne(ctx, newUser)
	if err != nil {
		log.Fatal("Erreur insertion :", err)
	}
	fmt.Println("Document added:", insertResult.InsertedID)

	filter := bson.D{{Key: "name", Value: "Melina Feutrier"}}
	var result bson.M
	err = collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal("Erreur lors de la recherche du document :", err)
	}
	fmt.Println("Document trouvé :", result)

	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal("Erreur lors de la recherche des documents :", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var doc bson.M
		if err := cursor.Decode(&doc); err != nil {
			log.Fatal("Erreur lors du décodage du document :", err)
		}
		fmt.Println("Document :", doc)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal("Erreur dans le curseur :", err)
	}
}
