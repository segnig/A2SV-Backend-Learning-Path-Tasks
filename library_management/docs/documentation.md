# 📚 Library Management System (Console-based) - GoLang

## 📝 Overview

The **Library Management System** is a simple console-based application built in Go to manage library books and members. It demonstrates the use of Go programming features such as **structs**, **interfaces**, **slices**, **maps**, and **methods**. It supports basic operations like adding, removing, borrowing, and returning books, along with member management.

---

## 🧱 Features

- Add and remove books from the library.
- Register new library members.
- Borrow books (if available).
- Return borrowed books.
- List all available books in the library.
- List all books borrowed by a specific member.
- Error handling for invalid operations (e.g., borrowing unavailable books).

---

## 🗂️ Folder Structure

```

library\_management/
├── main.go                          # Entry point of the application
├── controllers/
│   └── library\_controller.go       # Handles console inputs and user interaction
├── models/
│   ├── book.go                     # Defines Book struct
│   └── member.go                   # Defines Member struct
├── services/
│   └── library\_service.go          # Implements the LibraryManager interface and logic
├── docs/
│   └── documentation.md            # This documentation file
└── go.mod                          # Go module file

````

---

## 📦 Modules

### models/book.go

Defines the `Book` struct:
```go
type Book struct {
    ID     int
    Title  string
    Author string
    Status string // "Available" or "Borrowed"
}
````

---

### models/member.go

Defines the `Member` struct:

```go
type Member struct {
    ID            int
    Name          string
    BorrowedBooks []Book
}
```

---

### services/library\_service.go

Implements the `LibraryManager` interface:

```go
type LibraryManager interface {
    AddBook(book Book)
    RemoveBook(bookID int)
    BorrowBook(bookID int, memberID int) error
    ReturnBook(bookID int, memberID int) error
    ListAvailableBooks() []Book
    ListBorrowedBooks(memberID int) []Book
}
```

The `Library` struct is responsible for storing and manipulating book and member data.

---

### controllers/library\_controller.go

Provides a text-based console interface that allows the user to:

* Select operations from a menu
* Add/remove books
* Borrow/return books
* Register members
* View books and borrowed books

It uses standard input to read from the console and interacts with the service layer.

---

### main.go

Initializes the library and starts the console interface.

```go
func main() {
    library := services.NewLibrary()
    controllers.RunConsole(library)
}
```

---

## 🚀 How to Run the Application

### 1. Install Go (if not already)

Download and install from: [https://golang.org/dl/](https://golang.org/dl/)

### 2. Clone the project or create the folder structure manually

### 3. Initialize the Go module:

```sh
go mod init library_management
go mod tidy
```

### 4. Run the application:

```sh
go run main.go
```

---

## 🧪 Sample Usage

```
--- Library Management System ---
1. Add Book
2. Remove Book
3. Borrow Book
4. Return Book
5. List Available Books
6. List Borrowed Books by Member
7. Add Member
0. Exit
```

Example Flow:

1. Add Member (e.g., ID: 1, Name: Alice)
2. Add Book (e.g., ID: 100, Title: "Go Basics", Author: "John Doe")
3. Borrow Book (Book ID: 100, Member ID: 1)
4. List Borrowed Books by Member
5. Return Book

---

## 🔧 Error Handling

* ❌ Attempting to borrow a non-existent or already borrowed book will display an error.
* ❌ Returning a book not borrowed or by an invalid member will raise an error.
* ✅ All operations validate input IDs and statuses.

---

## 💡 Concepts Demonstrated

* **Structs** (`Book`, `Member`)
* **Interfaces** (`LibraryManager`)
* **Slices** (list of borrowed books)
* **Maps** (book and member storage)
* **Methods** (attached to `Library`)
* **Error handling**
* **Console I/O**

---

## 🔄 Future Enhancements

* Add persistent storage (e.g., save/load to JSON or database)
* Add date tracking for borrowed books
* Limit number of books a member can borrow
* Web or GUI interface using Go frameworks

---

## 👨‍💻 Author

Developed by: **\[Segni Girma]**
GitHub: \[[Library Management](https://github.com/segnig/A2SV-Backend-Learning-Path-Tasks/library_management)]
Email: \[[segnigirma11@gmail.com](mailto:segnigirma11@gmail.com)]

---

## 📄 License

This project is licensed under the MIT License. See the LICENSE file for details.