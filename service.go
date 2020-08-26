package main

type IBookRepository interface {
	FindById(id string) (*Book, error)
	Save(*Book) error
}

type BookService struct {
	repo IBookRepository
}

func (svc *BookService) GetBook(id string) (*Book, error) {

	book, err := svc.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (svc *BookService) ReviewBook(id string, point float64) error {

	book, err := svc.repo.FindById(id)
	if err != nil {
		return err
	}

	book.AddReview(point)

	err = svc.repo.Save(book)
	if err != nil {
		return err
	}

	return nil
}