package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Email    string `json:"email" bson:"email"`
}

func connectToMongo() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://MelinaFeu:123456..@cluster0.grdqi.mongodb.net/"))
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to Database")
	}

	return client
}

func getUserByID(client *mongo.Client, userID string) (*User, error) {
	collection := client.Database("testdb").Collection("users")
 
	// Convertir userID en ObjectID
	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID format")
	}
 
	var user User
	filter := bson.M{"_id": id} 
 
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}
 
	return &user, nil
}

// Nouvelle fonction pour récupérer tous les utilisateurs
func getAllUsers(client *mongo.Client) ([]User, error) {
	collection := client.Database("testdb").Collection("users")

	var users []User

	// Requête sans filtre pour récupérer tous les documents
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	// Parcourir les documents et les décoder
	for cursor.Next(context.Background()) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	client := connectToMongo()
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal(err)
		}
		log.Println("Disconnected from Database")
	}()

	r := gin.Default()

	r.GET("/user/:id", func(c *gin.Context) {
		userID := c.Param("id")

		user, err := getUserByID(client, userID)
		if err != nil {
			c.JSON(404, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
		})
	})

	// Nouvelle route pour récupérer tous les utilisateurs
	r.GET("/users", func(c *gin.Context) {
		users, err := getAllUsers(client)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, users)
	})

	r.Run(":8080")
}