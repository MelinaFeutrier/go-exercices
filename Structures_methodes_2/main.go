package main

import "fmt"

type Book struct {
	Title           string
	Author          string
	PublicationYear int
}

type Library struct {
	Books []Book
}

func (l *Library) AddBook(book Book) {
	l.Books = append(l.Books, book) 
	fmt.Printf("Livre ajouté : \"%s\" par %s (%d)\n", book.Title, book.Author, book.PublicationYear)
}

func (l Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("La bibliothèque est vide.")
		return
	}
	fmt.Println("Liste des livres dans la bibliothèque :")
	for i, book := range l.Books {
		fmt.Printf("%d. \"%s\" par %s (%d)\n", i+1, book.Title, book.Author, book.PublicationYear)
	}
}

func main() {
	myLibrary := Library{}

	myLibrary.AddBook(Book{Title: "1984", Author: "George Orwell", PublicationYear: 1949})
	myLibrary.AddBook(Book{Title: "Le Petit Prince", Author: "Antoine de Saint-Exupéry", PublicationYear: 1943})
	myLibrary.AddBook(Book{Title: "Les Misérables", Author: "Victor Hugo", PublicationYear: 1862})

	myLibrary.ListBooks()
}
