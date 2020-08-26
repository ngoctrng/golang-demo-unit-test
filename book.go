package main

import "math"

type Book struct {
	ID            string
	Title         string
	PointTotal    float64
	PointShow     float64
	TotalReviewer int
}

func (b *Book) AddReview(point float64) {
	b.TotalReviewer++
	b.PointTotal += point
	newPoint := b.PointTotal / float64(b.TotalReviewer)
	b.PointShow = math.Round(newPoint*10) / 10
}
