package main

import (
	"log"

	"github.com/JoaoRafa19/point-and-click/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("PAC")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game.NewGame()

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
