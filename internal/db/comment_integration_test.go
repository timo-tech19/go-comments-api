//go:build integration
// +build integration

package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/timo-tech19/go-comments-api/internal/comment"
	"testing"
)

func TestCommentDatabase(t *testing.T) {
	t.Run("test create comment database", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)

		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "author",
			Body:   "body",
		})
		assert.NoError(t, err)

		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		asser.Equal(t, "slug", newCmt.Slug)
		fmt.Println("Testing comment creation")
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug: "new-slug",
			Author: "Timo",
			Body: "body"
		})
		assert.NoError(t, err)

		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)

		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.Error(t, err)
	})
}
