package store

import (
	"fmt"

	"github.com/ea3hsp/book/pkg/models"
)

func (sr *simdbRepository) CreateAuth(auth *models.Author) (string, error) {
	// insert book
	err := sr.db.Insert(auth)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("created author ID: [%s]", auth.AuthorID), nil
}

func (sr *simdbRepository) GetAuthor(authID string) (*models.Author, error) {
	// Author holder
	auth := &models.Author{}
	// select Author id
	err := sr.db.Open(models.Author{}).Where("authid", "=", authID).First().AsEntity(&auth)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func (sr *simdbRepository) GetAuthors() (*[]models.Author, error) {
	// auths holder
	auths := &[]models.Author{}
	// select all auths
	err := sr.db.Open(models.Author{}).Get().AsEntity(&auths)
	if err != nil {
		return nil, err
	}
	return auths, nil
}
