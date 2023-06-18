package comment

import (
	"context"
	"fmt"
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	author string
}

// interfaces in go???
type Store interface {
	GetComment(context.Context, string) (Comment, error) // on the db level
}

// Service - struct on which all our logic will be built on
type Service struct {
	Store Store
}

// NewService - returns a pointer to a newly created service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - method on a service struct which contains service level logic
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}

	return cmt, nil
}
