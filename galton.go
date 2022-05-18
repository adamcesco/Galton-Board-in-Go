package main

import "github.com/adamcesco/Galton-Board-in-Go/GaltonStructs"

func main() {
	board, err := GaltonStructs.CreateGaltonBoard(30, 28, 3000)
	if err != nil {
		panic(err)
	}
	board.CalculateResults()
	board.PrintResults()
	board.DisplayResults()
}
