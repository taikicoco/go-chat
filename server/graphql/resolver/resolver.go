package resolver

import (
	"sync"

	"server/graphql/generated/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ChatID map[int64][]chan<- *model.Message
	Mutex  sync.Mutex
}
