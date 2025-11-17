package main

import (
	"fmt"
	"os"

	"github.com/avraam311/distributed-gogrep/internal/coordinator"
)

func main() {
	coord := coordinator.NewCoordinator()
	err := coord.StartGrepping()
	if err != nil {
		fmt.Println("failed to process text")
		os.Exit(1)
	}
}
