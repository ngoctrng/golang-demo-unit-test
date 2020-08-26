package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBook_AddReview(t *testing.T) {
	t.Run("Test Review a book", func(t *testing.T) {
		b := &Book{
			ID:            "test_id",
			Title:         "test title",
			PointTotal:    19,
			PointShow:     3.8,
			TotalReviewer: 5,
		}
		b.AddReview(4.9)

		assert.Equal(t, 4.0, b.PointShow)
	})
}
