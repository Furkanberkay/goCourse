// Package interfaces, domain: geometry primitives with a clean, safe API.
package interfaces

import (
	"errors"
	"fmt"
	"math"
)

// Shape = sözleşme. Saf ölçüm metodları içindir.
// Not: İsim export edildi (S büyük) ki paket dışından da tüketilebilsin.
type Shape interface {
	Area() float64
	// İleride Perimeter() gibi ek metrikler buraya eklenebilir (interface segregation'ı bozmadan).
}

// Rectangle: Alan/çevre hesabına uygun, immutability'ye yakın tasarım.
// Alanlar unexport (küçük harf) => dışarıdan serbest set edilemez; tek giriş kapısı constructor.
type Rectangle struct {
	width  float64
	height float64
}

// Circle: Aynı güvenlik modeli; tek giriş kapısı constructor.
type Circle struct {
	radius float64
}

// --- Constructors (Validation + Invariants) ---

// NewRectangle: domain kuralları burada enforce edilir.
func NewRectangle(width, height float64) (Rectangle, error) {
	if width <= 0 || height <= 0 {
		return Rectangle{}, errors.New("width and height must be > 0")
	}
	return Rectangle{width: width, height: height}, nil
}

// NewCircle: radius validasyonu.
func NewCircle(radius float64) (Circle, error) {
	if radius <= 0 {
		return Circle{}, errors.New("radius must be > 0")
	}
	return Circle{radius: radius}, nil
}

// --- Getters (read-only erişim) ---

func (r Rectangle) Width() float64  { return r.width }
func (r Rectangle) Height() float64 { return r.height }
func (c Circle) Radius() float64    { return c.radius }

// ------ Setters

func (r *Rectangle) SetRectangleWidth(w float64) error {
	if w <= 0 {
		return fmt.Errorf("width must be > 0")
	}
	r.width = w
	return nil
}

func (r *Rectangle) SetRectangleHeight(h float64) {
	r.height = h
}

func (r *Rectangle) SetRectangleHeightAndWidth(h, w float64) error {
	if h <= 0 || w <= 0 {
		return fmt.Errorf("height and width must be > 0")
	}
	r.height = h
	r.width = w
	return nil
}

// --- Pure methods (no side effects) ---

// Value receiver: küçük, immutable veri; kopya maliyeti düşük.
func (r Rectangle) Area() float64 { return r.width * r.height }

func (c Circle) Area() float64 { return math.Pi * c.radius * c.radius }

// --- Application boundary (demo/use-case) ---

// PrintArea: Interface üzerinden çalışır; polimorfik ve test edilebilir.
func PrintArea(s Shape) float64 {
	area := s.Area()
	fmt.Println("Alan:", area)
	return area
}
