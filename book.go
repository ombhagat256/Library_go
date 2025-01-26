package main

import "fmt"

// Book struct
type Book struct {
	Title     string
	Author    string
	ISBN      string
	Available bool
}

// Method to initialize a new book
func NewBook(title, author, isbn string, available bool) Book {
	return Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: available,
	}
}

// Method to display book details
func (b Book) DisplayDetails() {
	fmt.Printf("Title: %s, Author: %s, ISBN: %s, Available: %t\n", b.Title, b.Author, b.ISBN, b.Available)
}

// EBook struct embedding Book
type EBook struct {
	Book
	FileSize int // File size in MB
}

// Overriding the display method to include file size
func (e EBook) DisplayDetails() {
	fmt.Printf("Title: %s, Author: %s, ISBN: %s, Available: %t, File Size: %dMB\n", e.Title, e.Author, e.ISBN, e.Available, e.FileSize)
}

// BookInterface
type BookInterface interface {
	DisplayDetails()
}
