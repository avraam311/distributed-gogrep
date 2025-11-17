package coordinator

import "fmt"

type Agregator struct {
	matches [][]string
}

func NewAgregator(matches [][]string) *Agregator {
	return &Agregator{}
}

func (a *Agregator) PrintResultWithNumbers() {
	for _, pair := range a.matches {
		fmt.Print(pair[0])
		fmt.Println(pair[1])
	}
}

func (a *Agregator) PrintResult() {
	for _, match := range a.matches {
		fmt.Println(match[1])
	}
}
