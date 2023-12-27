package resolver

import (
	"server/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	user       *usecase.User
}

func NewResolver(user *usecase.User) *Resolver {
	return &Resolver{
		user:       user,	
	}
}
