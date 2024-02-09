package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"server/graphql/generated/model"
)

const postStatus = "sent"

// PostMessage is the resolver for the postMessage field.
func (r *mutationResolver) PostMessage(ctx context.Context, id int64, text string) (*model.Message, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	for _, ch := range r.MessageID[int64(id)] {
		ch <- &model.Message{
			ID:   id,
			Text: text,
		}
	}
	return &model.Message{
		ID:   id,
		Text: text,
		Type: postStatus,
	}, nil
}

// Messages is the resolver for the messages field.
func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	return []*model.Message{
		{
			ID:   1,
			Text: "Hello World",
		},
	}, nil
}

// MessagePosted is the resolver for the messagePosted field.
func (r *subscriptionResolver) MessagePosted(ctx context.Context, id int64) (<-chan *model.Message, error) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	ch := make(chan *model.Message, 1)
	r.MessageID[id] = append(r.MessageID[id], ch)

	go func() {
		<-ctx.Done()
		r.Mutex.Lock()
		defer r.Mutex.Unlock()
		for i, c := range r.MessageID[id] {
			if c == ch {
				r.MessageID[id] = append(r.MessageID[id][:i], r.MessageID[id][i+1:]...)
				break
			}
		}
	}()

	return ch, nil
}
