package services

import (
	"errors"  // added package for error handling
	"strings" // new import for search functionality
	"sync"
	"test/models"
)

type BookStore struct {
	books []models.Book
	mu    sync.Mutex
}

var Store = BookStore{
	books: []models.Book{},
}

func (bs *BookStore) GetBooks() []models.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	return bs.books
}

func (bs *BookStore) AddBook(title, author string) models.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	newID := len(bs.books) + 1
	newBook := models.Book{ID: newID, Title: title, Author: author}
	bs.books = append(bs.books, newBook)
	return newBook
}

// New method to update an existing book.
func (bs *BookStore) UpdateBook(id int, title, author string) (models.Book, error) {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	for i, book := range bs.books {
		if book.ID == id {
			bs.books[i].Title = title
			bs.books[i].Author = author
			return bs.books[i], nil
		}
	}
	return models.Book{}, errors.New("book not found")
}

// New method to delete an existing book.
func (bs *BookStore) DeleteBook(id int) error {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	for i, book := range bs.books {
		if book.ID == id {
			bs.books = append(bs.books[:i], bs.books[i+1:]...)
			return nil
		}
	}
	return errors.New("book not found")
}

// New method to search books by title (case-insensitive).
func (bs *BookStore) SearchBooks(name string) []models.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	var results []models.Book
	lowerName := strings.ToLower(name)
	for _, book := range bs.books {
		if strings.Contains(strings.ToLower(book.Title), lowerName) {
			results = append(results, book)
		}
	}
	return results
}

// New method to search books by author (case-insensitive).
func (bs *BookStore) SearchBooksByAuthor(author string) []models.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()
	var results []models.Book
	lowerAuthor := strings.ToLower(author)
	for _, book := range bs.books {
		if strings.Contains(strings.ToLower(book.Author), lowerAuthor) {
			results = append(results, book)
		}
	}
	return results
}
