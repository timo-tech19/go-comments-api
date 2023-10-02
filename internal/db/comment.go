package db

import (
	"context"
	"database/sql"
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/timo-tech19/go-comments-api/internal/comment"
)

// CommentRow struct represents data from database.
// some fields are of type sql.NullString to handle potential null values in the database.
type CommentRow struct {
	ID     string
	Slug   sql.NullString
	Body   sql.NullString
	Author sql.NullString
}

// Converts CommentRow struct used in database layer to Comment struct used in service layer.
func commentRowToComment(c CommentRow) comment.Comment {
	return comment.Comment{
		ID:     c.ID,
		Slug:   c.Slug.String,
		Author: c.Author.String,
		Body:   c.Body.String,
	}
}

// Retrieves a comment from the database
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
		return comment.Comment{}, fmt.Errorf("error fetching comment by uuid, %w", err)
	}

	return commentRowToComment(cmtRow), nil
}

// Adds a new comment to the database.
func (d *Database) PostComment(ctx context.Context, cmt comment.Comment) (comment.Comment, error) {

	// incomming cmt.ID is "" so must be generated and assigned before sending to database
	cmt.ID = uuid.NewV4().String()
	postRow := CommentRow{
		ID:     cmt.ID,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	// Field values from postRow are inserted into the SQL query string.
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`INSERT INTO comments
		(id, slug, author, body)
		VALUES
		(:id, :slug, :author, :body)`,
		postRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("Failed to insert comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("Failed to close rows: %w", err)
	}

	// cmt struct is returned with cmt.ID updated. Returned data from database is not used here
	return cmt, nil
}

// Delete comment from database.
func (d *Database) DeleteComment(ctx context.Context, id string) error {
	_, err := d.Client.ExecContext(
		ctx,
		`DELETE FROM comments where id = $1`,
		id,
	)

	if err != nil {
		return fmt.Errorf("Failed to delete comment from database: %w", err)
	}

	return nil
}

// Updates comment in database
func (d *Database) UpdateComment(ctx context.Context, id string, cmt comment.Comment) (comment.Comment, error) {
	cmtRow := CommentRow{
		ID:     id,
		Slug:   sql.NullString{String: cmt.Slug, Valid: true},
		Author: sql.NullString{String: cmt.Author, Valid: true},
		Body:   sql.NullString{String: cmt.Body, Valid: true},
	}

	// fields from cmtRow are inserted into SQL query string
	rows, err := d.Client.NamedQueryContext(
		ctx,
		`UPDATE comments SET
		slug = :slug,
		author = :author,
		body = :body
		WHERE id = :id`,
		cmtRow,
	)

	if err != nil {
		return comment.Comment{}, fmt.Errorf("Failed to update comment: %w", err)
	}

	if err := rows.Close(); err != nil {
		return comment.Comment{}, fmt.Errorf("Failed to close rows: %w", err)
	}

	// A service comment struct is returned. Return value from database is not used here.
	return commentRowToComment(cmtRow), nil
}
