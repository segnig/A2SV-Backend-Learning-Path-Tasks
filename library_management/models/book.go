package model

type Book struct {
	ID     int
	Title  string
	Author string
	Status string // "Available" or "Borrowed"
}
