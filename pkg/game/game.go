package game

import (
	"math/rand"
)

const (
	CHOICE_FIRE  = "fire"
	CHOICE_WATER = "water"
	CHOICE_GRASS = "grass"

	WINNER_CPU    = "Cpu"
	WINNER_PLAYER = "Player"
)

type Game struct {
	PlayerChoice string
	CpuChoice    string
	GameOver     bool
	Draw         bool
	Decision     string
}

func New() *Game {
	g := Game{}
	return &g
}

func (g *Game) UserChoice(userInput string) {
	switch userInput {
	case "fire":
	case "f":
		g.PlayerChoice = CHOICE_FIRE
	case "water":
	case "w":
		g.PlayerChoice = CHOICE_WATER
	case "grass":
	case "g":
		g.PlayerChoice = CHOICE_GRASS
	default:
		panic("no choice discernable")
	}
}

func (g *Game) Choose() {
	choice := rand.Intn(3)
	switch choice {
	case 0:
		g.CpuChoice = CHOICE_GRASS
	case 1:
		g.CpuChoice = CHOICE_FIRE
	case 2:
		g.CpuChoice = CHOICE_WATER
	}
}

func (g *Game) HasChosen() bool {
	if g.CpuChoice != "" {
		return true
	}
	return false
}

func (g *Game) Decide() {
	if g.PlayerChoice == g.CpuChoice {
		g.Draw = true
		g.GameOver = true
		return
	}

	if g.PlayerChoice == CHOICE_FIRE {
		if g.CpuChoice == CHOICE_WATER {
			g.Decision = WINNER_CPU
		} else {
			g.Decision = WINNER_PLAYER
		}
	}
	if g.PlayerChoice == CHOICE_GRASS {
		if g.CpuChoice == CHOICE_WATER {
			g.Decision = WINNER_PLAYER
		} else {
			g.Decision = WINNER_CPU
		}
	}
  if g.PlayerChoice == CHOICE_WATER {
		if g.CpuChoice == CHOICE_FIRE {
			g.Decision = WINNER_PLAYER
		} else {
			g.Decision = WINNER_CPU
		}
	}
	g.GameOver = true
}

func (g *Game) IsEnded() bool {
	if g.GameOver {
		return true
	}
	return false
}
