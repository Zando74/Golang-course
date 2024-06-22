// Liskov Substitution Principle
package main

import "fmt"

/*
	LSP : Liskov Substitution Principle
	un objet de type T peut être remplacé par un objet de type S sans altérer les propriétés désirables du programme
	-> Un sous-type doit être substituable à son type de base
*/

type Sized interface {
	GetWidth() int
	SetWidth(width int)
	GetHeight() int
	SetHeight(height int)
}

type Rectangle struct {
	width, height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square {
	square := Square{}
	square.width = size
	square.height = size
	return &square
}

func (s *Square) SetWidth(width int) {
	s.width = width
	s.height = width
}

func (s *Square) SetHeight(height int) {
	s.width = height
	s.height = height
}

func useIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Println("Expected area: ", expectedArea, ", actual area: ", actualArea)
}

func main() {
	rc := &Rectangle{2, 3}
	useIt(rc)

	sq := NewSquare(5)
	useIt(sq)
}

// Pas d'héritage en Go, donc pas de solution simple pour le LSP
