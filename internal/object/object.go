package object

import (
	"fish/internal/point"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update() error
	Draw(screen *ebiten.Image)
	TouchesObject(Object) bool
	GetTouchArea() (*point.Point, float64)
	GetSize() int
	IsDead() bool
	Kill()
}
