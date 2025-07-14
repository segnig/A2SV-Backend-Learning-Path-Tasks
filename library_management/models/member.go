package model

type Member struct {
	ID            int
	Name          string
	BorrowedBooks []Book
}
