package store

import "github.com/ea3hsp/book/pkg/models"

// IStore interface
type IStore interface {
	// Library
	CreateBook(*models.Book) (string, error)
	DeleteBook(string) (string, error)
	GetBook(string) (*models.Book, error)
	GetBooks() (*[]models.Book, error)
	ModifyBook(*models.Book) (string, error)
	// Authors
	CreateAuth(*models.Author) (string, error)
	GetAuthor(string) (*models.Author, error)
	GetAuthors() (*[]models.Author, error)
	// Publishers
	CreatePublisher(*models.Publisher) (string, error)
	GetPublisher(string) (*models.Publisher, error)
	GetPublishers() (*[]models.Publisher, error)
}
