package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	library := Library{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLibrary Menu:")
		fmt.Println("1. Add a Book")
		fmt.Println("2. Add an EBook")
		fmt.Println("3. Remove a Book/EBook")
		fmt.Println("4. Search for Books by Title")
		fmt.Println("5. List All Books")
		fmt.Println("6. Exit")

		fmt.Print("Choose an option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			addBook(reader, &library)
		case "2":
			addEBook(reader, &library)
		case "3":
			removeBook(reader, &library)
		case "4":
			searchBook(reader, library)
		case "5":
			library.ListBooks()
		case "6":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func addBook(reader *bufio.Reader, library *Library) {
	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter ISBN: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	fmt.Print("Is Available (true/false): ")
	availableStr, _ := reader.ReadString('\n')
	available, _ := strconv.ParseBool(strings.TrimSpace(availableStr))

	book := NewBook(title, author, isbn, available)
	library.AddBook(book)
	fmt.Println("Book added successfully.")
}

func addEBook(reader *bufio.Reader, library *Library) {
	fmt.Print("Enter Title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Enter Author: ")
	author, _ := reader.ReadString('\n')
	author = strings.TrimSpace(author)

	fmt.Print("Enter ISBN: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	fmt.Print("Is Available (true/false): ")
	availableStr, _ := reader.ReadString('\n')
	available, _ := strconv.ParseBool(strings.TrimSpace(availableStr))

	fmt.Print("Enter File Size (MB): ")
	fileSizeStr, _ := reader.ReadString('\n')
	fileSize, _ := strconv.Atoi(strings.TrimSpace(fileSizeStr))

	ebook := EBook{
		Book:     NewBook(title, author, isbn, available),
		FileSize: fileSize,
	}
	library.AddBook(ebook)
	fmt.Println("EBook added successfully.")
}

func removeBook(reader *bufio.Reader, library *Library) {
	fmt.Print("Enter ISBN to remove: ")
	isbn, _ := reader.ReadString('\n')
	isbn = strings.TrimSpace(isbn)

	library.RemoveBook(isbn)
}

func searchBook(reader *bufio.Reader, library Library) {
	fmt.Print("Enter Title to search: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	library.SearchBookByTitle(title)
}
