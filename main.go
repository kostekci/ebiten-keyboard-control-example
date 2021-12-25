package main

import (
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("gopher.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Character struct {
	x float64
	y float64
}

func (c *Character) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		c.x -= 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		c.x += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		c.y += 5
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		c.y -= 5
	}
}

type Game struct {
	character Character
	inited    bool
}

func (g *Game) init() {
	defer func() {
		g.inited = true
	}()

	g.character.x = 0
	g.character.y = 0
}

func (g *Game) Update() error {
	if !g.inited {
		g.init()
	}
	g.character.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.character.x, g.character.y)

	screen.DrawImage(img, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Geometry Matrix")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
