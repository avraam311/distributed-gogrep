package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/avraam311/distributed-gogrep/internal/coordinator"
	"github.com/avraam311/distributed-gogrep/internal/grepper"
)

func main() {
	ports := []string{"8081", "8082", "8083", "8084", "8085"}
	var wg sync.WaitGroup

	for _, port := range ports {
		flags := &coordinator.Parser{}
		g := grepper.NewGrepper(flags.Strings, flags.Template)
		app := grepper.NewApp(g, flags)

		wg.Add(1)
		go grepper.StartServer(port, app, &wg)
	}

	wg.Wait()

	coord := coordinator.NewCoordinator()
	err := coord.StartGrepping()
	if err != nil {
		fmt.Println("failed to process text")
		os.Exit(1)
	}
}
