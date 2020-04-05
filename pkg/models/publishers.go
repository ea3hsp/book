package models

// Publisher publisher representation
type Publisher struct {
	PublisherID string `json:"pubid"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Address     string `json:"address"`
	Email       string `json:"email"`
	URL         string `json:"url"`
	Country     string `json:"country"`
}

// ID id for publisher registries
func (p Publisher) ID() (jsonField string, value interface{}) {
	value = p.PublisherID
	jsonField = "pubid"
	return
}
