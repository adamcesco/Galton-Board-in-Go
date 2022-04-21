package main

import "main.go/structs"

func main() {
	board := structs.Galtonboard(30, 28, 3000)
	board.CalculateResults()
	board.PrintResults()
	board.DisplayResults()
}
