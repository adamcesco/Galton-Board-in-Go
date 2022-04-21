package main

import "main.go/GaltonStructs"

func main() {
	board := GaltonStructs.GaltonBoard(30, 28, 3000)
	board.CalculateResults()
	board.PrintResults()
	board.DisplayResults()
}
