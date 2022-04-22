package main

import "github.com/adamcesco/Galton-Board-in-Go/GaltonStructs"

func main() {
	board := GaltonStructs.GaltonBoard(30, 28, 3000)
	board.CalculateResults()
	board.PrintResults()
	board.DisplayResults()
}
