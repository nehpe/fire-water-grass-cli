package main

import (
	"fmt"
	"os"

	"github.com/onwhee/go-fwg/pkg/game"
)

func main() {
  args := os.Args[1:]
  choice := args[0]
  fmt.Println("Choice: "+choice)

  g := game.New()
  g.UserChoice(choice)
  g.Choose()
  g.Decide()

  if g.Draw {
    fmt.Println("Your game ended in a draw.")
  } else {
    if g.Decision == game.WINNER_PLAYER {
      fmt.Printf("You won! (%s vs. %s)", g.PlayerChoice, g.CpuChoice)
    } else {
      fmt.Printf("You lost! (%s vs. %s)", g.PlayerChoice, g.CpuChoice)
    }
  }
  fmt.Println()


}
