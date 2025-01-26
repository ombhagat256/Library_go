package main

import "fmt"

// Library struct to manage a collection of books
type Library struct {
	Books []BookInterface
}

// AddBook adds a book to the library
func (l *Library) AddBook(book BookInterface) {
	l.Books = append(l.Books, book)
}

// RemoveBook removes a book from the library by ISBN
func (l *Library) RemoveBook(isbn string) {
	for i, book := range l.Books {
		switch v := book.(type) {
		case Book:
			if v.ISBN == isbn {
				l.Books = append(l.Books[:i], l.Books[i+1:]...)
				fmt.Println("Book removed successfully")
				return
			}
		case EBook:
			if v.ISBN == isbn {
				l.Books = append(l.Books[:i], l.Books[i+1:]...)
				fmt.Println("EBook removed successfully")
				return
			}
		}
	}
	fmt.Println("Book not found")
}

// SearchBookByTitle searches for books by title
func (l Library) SearchBookByTitle(title string) {
	found := false
	for _, book := range l.Books {
		switch v := book.(type) {
		case Book:
			if v.Title == title {
				v.DisplayDetails()
				found = true
			}
		case EBook:
			if v.Title == title {
				v.DisplayDetails()
				found = true
			}
		}
	}
	if !found {
		fmt.Println("No books found with the title:", title)
	}
}

// ListBooks lists all books in the library
func (l Library) ListBooks() {
	if len(l.Books) == 0 {
		fmt.Println("No books in the library")
		return
	}
	for _, book := range l.Books {
		book.DisplayDetails()
	}
}
