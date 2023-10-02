package comment

import (
	"context" // context.Context provides a central location for passing data between different layers of our app
	"errors"
	"fmt"
)

// Define common error messages
var (
	ErrFetchingComment = errors.New("Failed to fetch comment by id")
	ErrPostingComment  = errors.New("Failed to create comment")
	ErrUpdatingComment = errors.New("Failed to update comment")
	ErrDeletingComment = errors.New("Failed to delete comment")
	ErrNotImplemented  = errors.New("Not implemented")
)

// Comment struct represent data in comment service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store defines operations to be implemented by the database layer
type Store interface {
	GetComment(context.Context, string) (Comment, error) // on the db level
	PostComment(context.Context, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	UpdateComment(context.Context, string, Comment) (Comment, error)
}

// Service struct models our service functionality
type Service struct {
	Store Store
}

// Returns a new instance of our comment service.
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// Retrieves comment data in service layer
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}

	return cmt, nil
}

// Updates comment data in service layer
func (s *Service) UpdateComment(ctx context.Context, ID string, cmt Comment) (Comment, error) {
	updatedCmt, err := s.Store.UpdateComment(ctx, ID, cmt)

	if err != nil {
		return Comment{}, ErrUpdatingComment
	}
	return updatedCmt, nil
}

// Deletes comment data in service layer
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	err := s.Store.DeleteComment(ctx, id)

	if err != nil {
		return ErrDeletingComment
	}

	return nil
}

// Adds a comment in service layer
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)

	if err != nil {
		return Comment{}, ErrPostingComment
	}

	return insertedCmt, nil
}
