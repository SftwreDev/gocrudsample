package models

type Actors struct {
	ID uint `json:"id" gorm:"primary_key"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
