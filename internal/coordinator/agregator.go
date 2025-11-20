package coordinator

import (
	"fmt"
	"log"
)

type Agregator struct {
	matches [][]string
}

func NewAgregator(matches [][]string) *Agregator {
	return &Agregator{
		matches: matches,
	}
}

func (a *Agregator) PrintResultWithNumbers() {
	log.Printf("aggregator: found %d matches", len(a.matches))
	for _, m := range a.matches {
		fmt.Println(m[1])
	}
}
