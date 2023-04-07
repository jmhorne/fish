package game

import (
	"fish/internal/fish"
	"fish/internal/object"
	"fish/internal/player"
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	score             int
	player            object.Object
	fishes            []object.Object
	level             int
	pointsToNextLevel int
	running           bool
	playerWon         bool
	width, height     int
	rando             *rand.Rand
}

func New(width, height int) (*Game, error) {
	var err error
	seed := rand.NewSource(time.Now().UnixNano())

	g := new(Game)
	g.rando = rand.New(seed)
	g.score = 0
	g.level = 1
	g.pointsToNextLevel = g.level * 10
	g.width = width
	g.height = height
	g.player, err = player.New(g.level, 0, 0, width, height)
	g.running = true
	g.playerWon = false
	g.generateFishes()

	return g, err
}

func (g *Game) generateFishes() {

	g.fishes = make([]object.Object, 0)

	for i := 0; i < g.level*100; i++ {
		g.genFish(g.level)
	}

	for i := 0; i < g.level; i++ {
		g.genFish(g.level + 1)
	}

	g.genFish(g.level + 4)
}

func (g *Game) genFish(size int) {

	x := g.rando.Intn(g.width)
	y := g.rando.Intn(g.height)
	dx := float64(g.rando.Intn(2) + 1)
	dy := float64(g.rando.Intn(2) + 1)
	f, err := fish.New(size, float64(x), float64(y), g.width, g.height, dx, dy)
	if err == nil {
		g.fishes = append(g.fishes, f)
	}
}

func (g *Game) Update() error {
	if !g.running {
		return nil
	}

	g.player.Update()

	for i, f := range g.fishes {
		f.Update()
		if g.player.TouchesObject(f) {
			if g.player.IsDead() {
				g.score += f.GetSize()
				g.fishes = append(g.fishes[:i], g.fishes[i+1:]...)

				if g.rando.Intn(1000) == 0 {
					g.genFish(g.level + 1)
				} else {
					g.genFish(g.level)
				}

			} else {
				fmt.Println("player died")
				g.running = false
			}

		}
	}

	if g.score >= g.pointsToNextLevel {
		if g.playerWon = g.player.(*player.Player).Grow(); g.playerWon {
			g.running = false
		}
		g.pointsToNextLevel *= 2
		g.level++
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.running {
		screen.Fill(color.RGBA{0, 0, 0xff, 0xff})

		for _, f := range g.fishes {
			f.Draw(screen)
		}

		g.player.Draw(screen)

		ebitenutil.DebugPrint(screen, fmt.Sprintf("Score: %d Level: %d", g.score, g.level))

		return
	}

	screen.Fill(color.RGBA{0, 0, 0, 0xFF})

	if g.playerWon {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("YOU WON! High Score: %d", g.score))
	} else {
		ebitenutil.DebugPrint(screen, fmt.Sprintf("YOU DIED! High Score: %d", g.score))
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}
