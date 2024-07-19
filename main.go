package main

import (
	"fmt"
	"time"
)

type Game struct {
	rows  int
	cols  int
	state [][]int
	tick  time.Duration
}

func initGame(game *Game) {
	game.rows = 40
	game.cols = 300
	game.tick = 20
	game.state = make([][]int, game.rows)

	for row := range game.state {
		game.state[row] = make([]int, game.cols)
		for col := range game.state[row] {
			game.state[row][col] = 0
		}
	}
	game.state[20][15] = 1
	game.state[20][16] = 1
	game.state[21][14] = 1
	game.state[21][15] = 1
	game.state[22][15] = 1

}

var neighbors = [][]int{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func countNeighbors(x int, y int, game *Game) int {
	count := 0
	for _, neighbor := range neighbors {
		row := x + neighbor[0]
		col := y + neighbor[1]

		if row < 0 || row >= game.rows {
			continue
		}

		if col < 0 || col >= game.cols {
			continue
		}

		if (game.state[row][col]) == 1 {
			count++
		}
	}
	return count
}

func updateState(game *Game) {
	var newState = make([][]int, game.rows)
	for row := range game.state {
		newState[row] = make([]int, game.cols)
		for col := range game.state[row] {
			newState[row][col] = game.state[row][col]

			neighborCount := countNeighbors(row, col, game)

			/*
				Any live cell with fewer than two live neighbours dies, as if by underpopulation.
				Any live cell with two or three live neighbours lives on to the next generation.
				Any live cell with more than three live neighbours dies, as if by overpopulation.
				Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.
			*/

			if neighborCount < 2 {
				newState[row][col] = 0
			}

			if neighborCount >= 4 {
				newState[row][col] = 0
			}

			if neighborCount == 3 {
				newState[row][col] = 1
			}
		}
	}

	game.state = newState
}

func printState(game *Game) {
	fmt.Print("\033[H\033[2J")

	for row := range game.state {
		for col := range game.state[row] {
			if game.state[row][col] == 0 {
				fmt.Print(" ")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	var game Game
	initGame(&game)
	ticker := time.NewTicker(time.Duration(game.tick) * time.Millisecond)

	for {
		<-ticker.C

		printState(&game)
		updateState(&game)
	}
}
