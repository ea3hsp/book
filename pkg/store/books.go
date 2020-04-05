package store

import (
	"fmt"

	"github.com/ea3hsp/book/pkg/models"
)

func (sr *simdbRepository) CreateBook(book *models.Book) (string, error) {
	// insert book
	err := sr.db.Insert(book)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("created book ID: [%s]", book.BookID), nil
}

func (sr *simdbRepository) DeleteBook(bookID string) (string, error) {
	// book ID to delete
	toDel := &models.Book{
		BookID: bookID,
	}
	// deleting book
	err := sr.db.Delete(toDel)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("deleted book ID: [%s]", bookID), nil
}

func (sr *simdbRepository) GetBook(bookID string) (*models.Book, error) {
	// book holder
	book := &models.Book{}
	// select book id
	err := sr.db.Open(models.Book{}).Where("isbn", "=", bookID).First().AsEntity(&book)
	if err != nil {
		return nil, err
	}
	return book, nil
}

func (sr *simdbRepository) GetBooks() (*[]models.Book, error) {
	// book holder
	books := &[]models.Book{}
	// select book id
	err := sr.db.Open(models.Book{}).Get().AsEntity(&books)
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (sr *simdbRepository) ModifyBook(book *models.Book) (string, error) {
	// update book
	err := sr.db.Update(book)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("updated book ID: [%s]", book.BookID), nil
}
