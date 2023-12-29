package resolver

import (
	"server/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	message       *usecase.Message
}

func NewResolver(message *usecase.Message) *Resolver {
	return &Resolver{
		message:       message,
	}
}
