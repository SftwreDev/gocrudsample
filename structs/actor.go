package structs

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

type ActorInput struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}
