package models

import "time"

// Author author representation
type Author struct {
	AuthorID  string    `json:"authid"`
	Name      string    `json:"name"`
	LastName  string    `json:"lastname"`
	Picture   bool      `json:"picture"`
	Country   string    `json:"country"`
	BornDate  time.Time `json:"borndate"`
	DeathDate time.Time `json:"deathdate,omitempty"`
	Bio       string    `json:"bio,omitempty"`
}

//ID any struct that needs to persist should implement this function defined
//in Entity interface.
func (a Author) ID() (jsonField string, value interface{}) {
	value = a.AuthorID
	jsonField = "authid"
	return
}

// GetBornDate get born date in locales time format
func (a Author) GetBornDate() string {
	return a.BornDate.Format("02/01/2006")
}
