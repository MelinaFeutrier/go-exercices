//Création d'une API en Go pour manipuler des données MongoDB

package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Book représente un livre
type Book struct {
	ID     primitive.ObjectID
	Title  string
	Author string
	Year   int
	Tags   []string
}

// Library représente une bibliothèque
type Library struct {
	ID    primitive.ObjectID
	Name  string
	Books []Book
}

var bookCollection *mongo.Collection
var libraryCollection *mongo.Collection
var ctx context.Context

func main() {
	// Configure Gin
	route := gin.Default()

	// MongoDB connection
	uri := "mongodb+srv://MelinaFeu:123456..@cluster0.grdqi.mongodb.net/"
	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Erreur de connexion à MongoDB :", err)
	}

	bookCollection = client.Database("testdb").Collection("books")
	libraryCollection = client.Database("testdb").Collection("libraries")

	fmt.Println("Connexion réussie à MongoDB")

	// Routes pour les livres
	route.POST("/books", createBook)
	route.GET("/books", getAllBooks)
	route.GET("/book/:id", getBookByID)
	route.GET("/searchBook", searchBooks)

	// Routes pour les bibliothèques
	route.POST("/libraries", createLibrary)
	route.GET("/libraries", getAllLibraries)
	route.GET("/library/:id", getLibraryByID)

	// Lancement du serveur
	if err := route.Run(":8080"); err != nil {
		log.Fatal("Erreur lors du démarrage du serveur :", err)
	}
}

// --- GESTION DES LIVRES ---

func createBook(c *gin.Context) {
	var book Book

	// Validation des données
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erreur": err.Error()})
		return
	}

	// Génération d'un nouvel ID
	book.ID = primitive.NewObjectID()

	// Insertion dans MongoDB
	result, err := bookCollection.InsertOne(ctx, book)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'insertion du livre"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

func getAllBooks(c *gin.Context) {
	cursor, err := bookCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des livres"})
		return
	}
	defer cursor.Close(ctx)

	var books []Book
	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage d'un livre"})
			return
		}
		books = append(books, book)
	}

	c.JSON(http.StatusOK, books)
}

func searchBooks(c *gin.Context) {
	// Récupérer les paramètres de la requête
	title := c.Query("title") // Paramètre "title" (facultatif)
	author := c.Query("author") // Paramètre "author" (facultatif)
	tag := c.Query("tag") // Paramètre "tag" (facultatif)

	// Construire le filtre dynamique
	filter := bson.M{}

	if title != "" {
		filter["title"] = bson.M{"$regex": title, "$options": "i"} // Recherche insensible à la casse
	}
	if author != "" {
		filter["author"] = bson.M{"$regex": author, "$options": "i"} // Recherche insensible à la casse
	}
	if tag != "" {
		filter["tags"] = tag // Match exact sur un tag
	}

	// Récupérer les livres de la collection
	cursor, err := bookCollection.Find(ctx, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des livres", "details": err.Error()})
		return
	}
	defer cursor.Close(ctx)

	// Lire les résultats dans une liste
	var books []Book
	for cursor.Next(ctx) {
		var book Book
		if err := cursor.Decode(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage des livres", "details": err.Error()})
			return
		}
		books = append(books, book)
	}

	// Vérifier les erreurs du curseur
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur avec le curseur", "details": err.Error()})
		return
	}

	// Retourner les livres trouvés
	c.JSON(http.StatusOK, books)
}


func getBookByID(c *gin.Context) {
	bookID := c.Param("id")

	// Conversion de l'ID en ObjectID
	id, err := primitive.ObjectIDFromHex(bookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	// Recherche du livre
	var book Book
	err = bookCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Livre non trouvé"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération du livre"})
		return
	}

	c.JSON(http.StatusOK, book)
}

// --- GESTION DES BIBLIOTHÈQUES ---

func createLibrary(c *gin.Context) {
	var library Library

	// Validation des données envoyées
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Données invalides", "details": err.Error()})
		log.Println("Erreur de validation :", err)
		return
	}

	// Initialisation de l'ID si nécessaire
	library.ID = primitive.NewObjectID()

	// Validation et initialisation des livres
	for i := range library.Books {
		if library.Books[i].ID.IsZero() {
			library.Books[i].ID = primitive.NewObjectID()
		}
		if library.Books[i].Title == "" || library.Books[i].Author == "" || library.Books[i].Year == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Chaque livre doit avoir un titre, un auteur et une année"})
			return
		}
	}

	// Insertion dans MongoDB
	result, err := libraryCollection.InsertOne(ctx, library)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'insertion de la bibliothèque"})
		log.Println("Erreur MongoDB :", err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
}

func getAllLibraries(c *gin.Context) {
	cursor, err := libraryCollection.Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération des bibliothèques"})
		return
	}
	defer cursor.Close(ctx)

	var libraries []Library
	for cursor.Next(ctx) {
		var library Library
		if err := cursor.Decode(&library); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du décodage d'une bibliothèque"})
			return
		}
		libraries = append(libraries, library)
	}

	c.JSON(http.StatusOK, libraries)
}

func getLibraryByID(c *gin.Context) {
	libraryID := c.Param("id")

	// Conversion de l'ID en ObjectID
	id, err := primitive.ObjectIDFromHex(libraryID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID invalide"})
		return
	}

	// Recherche de la bibliothèque
	var library Library
	err = libraryCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&library)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, gin.H{"error": "Bibliothèque non trouvée"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la récupération de la bibliothèque"})
		return
	}

	c.JSON(http.StatusOK, library)
}