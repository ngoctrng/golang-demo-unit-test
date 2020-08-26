package main

var (
	demoBook = Book{
		ID:            "123",
		Title:         "My Book",
		PointTotal:    19,
		PointShow:     3.8,
		TotalReviewer: 5,
	}
)

type MockRepository struct{}

func (r *MockRepository) FindById(id string) (*Book, error) {
	return &demoBook, nil
}

func (r *MockRepository) Save(*Book) error {
	return nil
}
