package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookService_GetBook(t *testing.T) {
	tests := []struct {
		name string
		repo IBookRepository
		id   string
		want *Book
	}{
		{
			name: "Test get book by id",
			repo: &MockRepository{},
			id:   demoBook.ID,
			want: &demoBook,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &BookService{
				repo: tt.repo,
			}

			got, err := svc.GetBook(tt.id)

			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestBookService_ReviewBook(t *testing.T) {
	type args struct {
		id    string
		point float64
	}
	tests := []struct {
		name string
		repo IBookRepository
		args args
	}{
		{
			name: "Test Adding a review",
			repo: &MockRepository{},
			args: args{id: demoBook.ID, point: 4.5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &BookService{
				repo: tt.repo,
			}

			err := svc.ReviewBook(tt.args.id, tt.args.point)

			assert.Nil(t, err)
			assert.Equal(t, demoBook.PointShow, 3.9)
		})
	}
}
