package player

import (
	"fish/internal/object"
	"fish/internal/point"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var images = map[int]string{
	1: "internal/assets/player/player1.png",
	2: "internal/assets/player/player2.png",
	3: "internal/assets/player/player3.png",
	4: "internal/assets/player/player4.png",
	5: "internal/assets/player/player5.png",
	6: "internal/assets/player/player6.png",
}

type Player struct {
	Size    int
	xPos    float64
	yPos    float64
	xMax    int
	yMax    int
	speed   float64
	image   *ebiten.Image
	isAlive bool
}

func New(size int, xPos, yPos float64, xMax, yMax int) (*Player, error) {
	var err error
	p := new(Player)

	p.Size = size
	p.xPos = xPos
	p.yPos = yPos
	p.xMax = xMax
	p.yMax = yMax
	p.speed = 2
	p.isAlive = true

	p.image, _, err = ebitenutil.NewImageFromFile(images[size])

	return p, err
}

func (p *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.yPos -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.yPos += p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.xPos -= p.speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.xPos += p.speed
	}

	// check bounds
	if p.xPos < 0 {
		p.xPos = float64(p.xMax)
	}
	if p.xPos > float64(p.xMax) {
		p.xPos = 0
	}
	if p.yPos < 0 {
		p.yPos = float64(p.yMax)
	}
	if p.yPos > float64(p.yMax) {
		p.yPos = 0
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.xPos, p.yPos)
	screen.DrawImage(p.image, op)
}

func (p *Player) TouchesObject(o object.Object) bool {
	pO, pR := p.GetTouchArea()
	oO, oR := o.GetTouchArea()

	distance := pO.GetDistance(oO)

	touches := distance <= pR || distance <= oR

	if touches {
		if o.GetSize() > p.Size {
			p.Kill()
		} else {
			o.Kill()
		}
	}

	return touches
}

func (p *Player) GetTouchArea() (*point.Point, float64) {
	midX := float64(p.image.Bounds().Dx() / 2)
	midY := float64(p.image.Bounds().Dy() / 2)
	originX := p.xPos + midX
	originY := p.yPos + midY

	origin := point.New(originX, originY)
	radius := midX

	if midY < midX {
		radius = midY
	}

	return origin, radius
}

func (p *Player) GetSize() int {
	return p.Size
}

func (p *Player) IsDead() bool {
	return !p.isAlive
}

func (p *Player) Kill() {
	p.isAlive = false
}

func (p *Player) Grow() bool {
	p.Size++

	if p.Size > 6 {
		return true
	}

	p.image, _, _ = ebitenutil.NewImageFromFile(images[p.Size])

	p.speed++
	return false
}