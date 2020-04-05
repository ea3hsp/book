package models

import "time"

// Book book representation
type Book struct {
	BookID      string    `json:"isbn"`
	AuthorID    string    `json:"authid,omitempty"`
	PublisherID string    `json:"pubid"`
	HasCover    bool      `json:"hascover"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Publisher   string    `json:"publisher"`
	PublishDate time.Time `json:"publishdate"`
	BoughtDate  time.Time `json:"boughtdate"`
	Readed      bool      `json:"readed"`
	ReadedDate  time.Time `json:"readeddate"`
}

//ID any struct that needs to persist should implement this function defined
//in Entity interface.
func (b Book) ID() (jsonField string, value interface{}) {
	value = b.BookID
	jsonField = "isbn"
	return
}
