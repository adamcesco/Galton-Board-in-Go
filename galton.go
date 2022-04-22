package main

import "github.com/adamcesco/Galton-Board-in-Go/GaltonStructs"

func main() {
	board := GaltonStructs.CreateGaltonBoard(30, 28, 3000)
	board.CalculateResults()
	board.PrintResults()
	board.DisplayResults()
}
