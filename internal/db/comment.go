package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/timo-tech19/go-comments-api/internal/comment"
)

// This approach allows for managing potential null values on the repository layer only
type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

func commentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

func (d *Database) GetComment(ctx context.Context, uuid string) (comment.Comment, error) {
	var cmtRow CommentRow
	row := d.Client.QueryRowContext(
		ctx,
		`SELECT id, slug, body, author
		FROM comments
		WHERE id = $1`,
		uuid,
	)

	err := row.Scan(&cmtRow.ID, &cmtRow.Slug, &cmtRow.Body, &cmtRow.Author)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("Error fetching comment by uuid, %w", err)
	}

	return commentRowToComment(cmtRow), nil
}
