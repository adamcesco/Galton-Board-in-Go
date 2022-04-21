package main

import "main.go/structs"

const height = 12
const containerSize = 28
const numOfBeads = 3000

func main() {

	board := structs.Galtonboard(height, containerSize, numOfBeads)
	board.CalculateResults()
	board.PrintResults()

}
