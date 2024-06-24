package main

import "fmt"

type Renderer interface {
	RenderCircle(radius float32)
	RenderSquare(size float32)
}

type VectorRenderer struct {
}

func (v *VectorRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing a circle of radius", radius)
}

func (v *VectorRenderer) RenderSquare(size float32) {
	fmt.Println("Drawing a square of size", size)
}

type RasterRenderer struct {
	Dpi int
}

func (r *RasterRenderer) RenderCircle(radius float32) {
	fmt.Println("Drawing pixels for circle of radius", radius)
}

func (v *RasterRenderer) RenderSquare(size float32) {
	fmt.Println("Drawing pixels for square of size", size)
}

type Circle struct {
	renderer Renderer
	radius   float32
}

type Square struct {
	renderer Renderer
	size     float32
}

func (c *Circle) Draw() {
	c.renderer.RenderCircle(c.radius)
}

func NewCircle(renderer Renderer, radius float32) *Circle {
	return &Circle{renderer: renderer, radius: radius}
}

func NewSquare(renderer Renderer, size float32) *Square {
	return &Square{renderer: renderer, size: size}
}

func (c *Circle) Resize(factor float32) {
	c.radius *= factor
}

func (s *Square) Draw() {
	s.renderer.RenderSquare(s.size)
}

func main() {
	raster := RasterRenderer{}
	vector := VectorRenderer{}
	circle := NewCircle(&vector, 5)
	square := NewSquare(&raster, 10)
	circle.Draw()
	square.Draw()
}
