package structs

import "fmt"

type galtonBoard struct {
	layers        int
	containerSize int
	numOfBeads    int
	results       []int
	beads         []bead
}

func Galtonboard(layers, containerSize, numOfBeads int) galtonBoard {
	toReturn := galtonBoard{
		layers:        layers,
		containerSize: containerSize,
		numOfBeads:    numOfBeads,
		results:       make([]int, containerSize),
		beads:         make([]bead, numOfBeads),
	}

	for i := 0; i < toReturn.numOfBeads; i++ {
		toReturn.beads[i].reset(toReturn.containerSize, toReturn.layers)
	}

	return toReturn
}

func (g *galtonBoard) Reset() {
	g.results = make([]int, g.containerSize)
	for i := 0; i < g.numOfBeads; i++ {
		g.beads[i].reset(g.containerSize, g.layers)
	}
}

func (g *galtonBoard) CalculateResults() {
	for i := 0; i < g.layers; i++ {
		for j := 0; j < g.numOfBeads; j++ {
			g.beads[j].advance()
		}
	}

	for _, bead := range g.beads {
		g.results[bead.x]++
	}
}

func (g *galtonBoard) PrintResults() {
	for _, it := range g.results {
		fmt.Print(it, " ")
	}
	fmt.Println()
}

func (g *galtonBoard) DisplayResults() {
	var max int
	buffer := g.numOfBeads / 120
	rCopy := g.results
	for i := 0; i < g.containerSize; i++ {
		if rCopy[i] < buffer && rCopy[i] > 0 {
			rCopy[i] = 1
		} else {
			rCopy[i] /= buffer
		}
		if rCopy[i] > max {
			max = rCopy[i]
		}
	}

	for max > 0 {
		for i := 0; i < g.containerSize; i++ {
			if rCopy[i] == max {
				fmt.Print("|*")
				rCopy[i]--
			} else {
				fmt.Print("| ")
			}
		}
		fmt.Println("|")

		max--
	}
}
