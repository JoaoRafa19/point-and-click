package game

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(_, _ int) (screenWidth, screenHeight int) {
	return 320, 240
}

func NewGame() *Game {
	return &Game{}
}

var _ ebiten.Game = (*Game)(nil)
