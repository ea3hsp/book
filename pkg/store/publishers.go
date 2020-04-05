package store

import (
	"fmt"

	"github.com/ea3hsp/book/pkg/models"
)

func (sr *simdbRepository) CreatePublisher(publisher *models.Publisher) (string, error) {
	// insert publisher
	err := sr.db.Insert(publisher)
	if err != nil {
		return err.Error(), err
	}
	return fmt.Sprintf("created publisher ID: [%s]", publisher.PublisherID), nil
}

func (sr *simdbRepository) GetPublisher(pubid string) (*models.Publisher, error) {
	// pub holder
	pub := &models.Publisher{}
	// select auth id
	err := sr.db.Open(models.Publisher{}).Where("pubid", "=", pubid).First().AsEntity(&pub)
	if err != nil {
		return nil, err
	}
	return pub, nil
}

func (sr *simdbRepository) GetPublishers() (*[]models.Publisher, error) {
	// pubs holder
	pubs := &[]models.Publisher{}
	// select publishers
	err := sr.db.Open(models.Book{}).Get().AsEntity(&pubs)
	if err != nil {
		return nil, err
	}
	return pubs, nil
}
