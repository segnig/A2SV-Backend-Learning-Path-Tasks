package controllers

import (
	"bufio"
	"fmt"
	model "library_management/models"
	"library_management/services"
	"os"
	"strconv"
	"strings"
)

func RunConsole(library *services.Library) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Library Management System ---")
		fmt.Println("1. Add Book")
		fmt.Println("2. Remove Book")
		fmt.Println("3. Borrow Book")
		fmt.Println("4. Return Book")
		fmt.Println("5. List Available Books")
		fmt.Println("6. List Borrowed Books by Member")
		fmt.Println("7. Add Member")
		fmt.Println("0. Exit")
		fmt.Print("Select option: ")

		optionStr, _ := reader.ReadString('\n')
		option, _ := strconv.Atoi(strings.TrimSpace(optionStr))

		switch option {
		case 1:
			fmt.Print("Enter Book ID: ")
			id := readInt(reader)
			fmt.Print("Enter Title: ")
			title := readLine(reader)
			fmt.Print("Enter Author: ")
			author := readLine(reader)
			library.AddBook(model.Book{ID: id, Title: title, Author: author})
			fmt.Println("Book added.")
		case 2:
			fmt.Print("Enter Book ID to remove: ")
			id := readInt(reader)
			library.RemoveBook(id)
			fmt.Println("Book removed.")
		case 3:
			fmt.Print("Enter Book ID to borrow: ")
			bookID := readInt(reader)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			if err := library.BorrowBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book borrowed.")
			}
		case 4:
			fmt.Print("Enter Book ID to return: ")
			bookID := readInt(reader)
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			if err := library.ReturnBook(bookID, memberID); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Book returned.")
			}
		case 5:
			fmt.Println("Available Books:")
			for _, b := range library.ListAvailableBooks() {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}
		case 6:
			fmt.Print("Enter Member ID: ")
			memberID := readInt(reader)
			fmt.Println("Borrowed Books:")
			for _, b := range library.ListBorrowedBooks(memberID) {
				fmt.Printf("ID: %d | Title: %s | Author: %s\n", b.ID, b.Title, b.Author)
			}
		case 7:
			fmt.Print("Enter Member ID: ")
			id := readInt(reader)
			fmt.Print("Enter Member Name: ")
			name := readLine(reader)
			library.Members[id] = model.Member{ID: id, Name: name, BorrowedBooks: make([]model.Book, 0)}
			fmt.Println("Member added.")
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}
}

func readInt(reader *bufio.Reader) int {
	text, _ := reader.ReadString('\n')
	num, _ := strconv.Atoi(strings.TrimSpace(text))
	return num
}

func readLine(reader *bufio.Reader) string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}
