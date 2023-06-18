package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed to fetch comment by id")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	author string
}

// interfaces in go???
// Store - interface defines all of the methods which our service needs to operate
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
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt *Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt *Comment) (Comment, error) {
	return Comment{}, ErrNotImplemented
}
