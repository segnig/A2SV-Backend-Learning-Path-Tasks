package services

import (
	"errors"
	model "library_management/models"
)

type LibraryManagement interface {
	AddBook(book model.Book)
	RemoveBook(bookID int)
	BorrowBook(bookID, memberID int) error
	ReturnBook(bookID, memberID int) error
	ListAvailableBooks() []model.Book
	ListBorrowedBooks(memberID int) []model.Book
}

type Library struct {
	Books   map[int]model.Book
	Members map[int]model.Member
}

func NewLibrary() *Library {
	return &Library{
		Books:   make(map[int]model.Book),
		Members: make(map[int]model.Member),
	}
}

func (lib *Library) AddBook(book model.Book) {
	book.Status = "Available"
	lib.Books[book.ID] = book
}

func (lib *Library) RemoveBook(bookID int) {
	delete(lib.Books, bookID)
}

func (lib *Library) BorrowBook(bookID, membetID int) error {
	book, exists := lib.Books[bookID]

	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Borrowed" {
		return errors.New("book is already borrowed")
	}

	member, exists := lib.Members[membetID]

	if !exists {
		return errors.New("member not found")
	}

	book.Status = "Borrowed"

	lib.Books[bookID] = book

	member.BorrowedBooks = append(member.BorrowedBooks, book)

	return nil
}

func (lib *Library) ReturnBook(bookID, memberID int) error {
	member, exists := lib.Members[memberID]

	if !exists {
		return errors.New("member not found")
	}

	book, exists := lib.Books[bookID]

	if !exists {
		return errors.New("book not found")
	}

	if book.Status == "Available" {
		return errors.New("book is not currently borrowed")
	}

	for idx, bk := range member.BorrowedBooks {
		if bk.ID == bookID {
			member.BorrowedBooks = append(member.BorrowedBooks[:idx], member.BorrowedBooks[idx+1:]...)
			break
		}
	}

	book.Status = "Available"
	lib.Books[bookID] = book

	return nil
}

func (lib *Library) ListAvailableBooks() []model.Book {
	var availableBook []model.Book

	for _, book := range lib.Books {
		if book.Status == "Available" {
			availableBook = append(availableBook, book)
		}
	}

	return availableBook
}

func (lib *Library) ListBorrowedBooks(memberID int) []model.Book {
	if member, exists := lib.Members[memberID]; exists {
		return member.BorrowedBooks
	}

	return nil
}
