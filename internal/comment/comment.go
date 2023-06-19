package comment

import (
	"context" // context.Context provides a central location for passing data between different layers of our app
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("Failed to fetch comment by id")
	ErrPostingComment  = errors.New("Failed to create comment")
	ErrUpdatingComment = errors.New("Failed to update comment")
	ErrDeletingComment = errors.New("Failed to delete comment")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// interfaces in go???
// Store - interface defines all of the methods which our service needs to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error) // on the db level
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
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

func (s *Service) UpdateComment(ctx context.Context, ID string, cmt Comment) (Comment, error) {
	updatedCmt, err := s.Store.UpdateComment(ctx, ID, cmt)

	if err != nil {
		return Comment{}, ErrUpdatingComment
	}
	return updatedCmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)

	if err != nil {
		return ErrDeletingComment
	}

	return nil
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		return Comment{}, ErrPostingComment
	}

	return insertedCmt, nil
}
