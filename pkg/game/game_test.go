package game_test

import (
	"testing"

	"github.com/onwhee/go-fwg/pkg/game"
)

func TestChoose(t *testing.T) {
  g := game.New()
  g.Choose()

}

func TestHasChosen(t *testing.T) {
  g := game.New()
  if (g.HasChosen()) {
    t.Fatalf("Game believes there has been a CPU choice, prior to there being a request for it.")
  }
  g.Choose()
  if (!g.HasChosen()) {
    t.Fatalf("Game believes no choice has been made, even though it was already requested.")
  }
}

func TestUserChoices(t *testing.T) {
  g := game.New()
  g.UserChoice("f")
  if g.PlayerChoice != game.CHOICE_FIRE {
    t.Fatalf("User chose fire, but got %s", g.PlayerChoice)
  }
  g.UserChoice("fire")
  if g.PlayerChoice != game.CHOICE_FIRE {
    t.Fatalf("User chose fire, but got %s", g.PlayerChoice)
  }
  g.UserChoice("w")
  if g.PlayerChoice != game.CHOICE_WATER {
    t.Fatalf("User chose water, but got %s", g.PlayerChoice)
  }
  g.UserChoice("water")
  if g.PlayerChoice != game.CHOICE_WATER {
    t.Fatalf("User chose water, but got %s", g.PlayerChoice)
  }
  g.UserChoice("g")
  if g.PlayerChoice != game.CHOICE_GRASS {
    t.Fatalf("User chose grass, but got %s", g.PlayerChoice)
  }
  g.UserChoice("grass")
  if g.PlayerChoice != game.CHOICE_GRASS {
    t.Fatalf("User chose grass, but got %s", g.PlayerChoice)
  }
}

func TestDraw(t *testing.T) {
  g := game.New()
  g.Choose()
  g.PlayerChoice = g.CpuChoice
  g.Decide()
  if g.Draw != true {
    t.Fatalf("CPU and Player chose the same option but it did not result in a draw.")
  }
}

func TestGameRules(t *testing.T) {
  g := game.New()
  g.PlayerChoice = game.CHOICE_FIRE
  g.CpuChoice = game.CHOICE_GRASS
  g.Decide()
  if g.Decision != game.WINNER_PLAYER {
    t.Fatalf("Player should've won (fire vs. grass), but did not")
  }
  g.CpuChoice = game.CHOICE_WATER
  g.Decide()
  if g.Decision != game.WINNER_CPU {
    t.Fatalf("CPU should've won (fire vs. water), but did not")
  }
}
