package main

import (
	"fmt"
	"main.go/structs"
)

const height = 12
const containerSize = 28
const numOfBeads = 3000

func main() {
	var results [28]int
	var contestants [numOfBeads]structs.Bead

	for i := 0; i < numOfBeads; i++ {
		contestants[i].Reset(containerSize, height)
	}

	for i := 0; i < height; i++ {
		for j := 0; j < numOfBeads; j++ {
			contestants[j].Advance()
		}
	}

	for _, contestant := range contestants {
		results[contestant.ReadX()]++
	}

	for _, it := range results {
		fmt.Print(it, " ")
	}
}
