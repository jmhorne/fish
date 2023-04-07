package fish

import (
	"fish/internal/object"
	"fish/internal/point"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images = map[int]string{
	1: "internal/assets/fish/fish1.png",
	2: "internal/assets/fish/fish2.png",
	3: "internal/assets/fish/fish3.png",
	4: "internal/assets/fish/fish4.png",
	5: "internal/assets/fish/fish5.png",
	6: "internal/assets/fish/fish6.png",
}

type Fish struct {
	Size    int
	xPos    float64
	yPos    float64
	xMax    int
	yMax    int
	dX, dY  float64
	image   *ebiten.Image
	isAlive bool
}

func New(size int, xPos, yPos float64, xMax, yMax int, dX, dY float64) (*Fish, error) {
	var err error
	f := new(Fish)

	f.Size = size
	f.xPos = xPos
	f.yPos = yPos
	f.xMax = xMax
	f.yMax = yMax
	f.dX = dX
	f.dY = dY
	f.isAlive = true
	f.image, _, err = ebitenutil.NewImageFromFile(images[size])

	return f, err
}

func (f *Fish) Update() error {
	f.xPos += f.dX
	f.yPos += f.dY

	// check bounds
	if f.xPos < 0 || f.xPos >= float64(f.xMax) {
		f.dX *= -1
	}
	if f.yPos < 0 || f.yPos > float64(f.yMax) {
		f.dY *= -1
	}
	return nil
}

func (f *Fish) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(f.xPos, f.yPos)
	screen.DrawImage(f.image, op)
}

func (f *Fish) TouchesObject(object.Object) bool {
	return false
}

func (f *Fish) GetTouchArea() (*point.Point, float64) {
	midX := float64(f.image.Bounds().Dx() / 2)
	midY := float64(f.image.Bounds().Dy() / 2)
	originX := f.xPos + midX
	originY := f.yPos + midY

	origin := point.New(originX, originY)
	radius := midX

	if midY < midX {
		radius = midY
	}

	return origin, radius
}

func (f *Fish) GetSize() int {
	return f.Size
}

func (f *Fish) IsDead() bool {
	return !f.isAlive
}

func (f *Fish) Kill() {
	f.isAlive = false
}
