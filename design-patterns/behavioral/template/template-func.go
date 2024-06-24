package main

import "fmt"

func PlayGame2(start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int) {
	start()
	for !haveWinner() {
		takeTurn()
	}
	fmt.Printf("Player %d wins.\n", winningPlayer())
}

func main2() {
	turn, maxTurns, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting a game of chess.")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n",
			turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurns
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame2(start, takeTurn, haveWinner, winningPlayer)
}
