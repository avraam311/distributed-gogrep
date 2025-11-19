package main

import (
	"log"
	"os"
	"time"

	"github.com/avraam311/distributed-gogrep/internal/coordinator"
	"github.com/avraam311/distributed-gogrep/pkg/grepper"
)

func main() {
	ports := [5]string{"8080", "8081", "8082", "8083", "8084"}
	go grepper.Run(ports)
	time.Sleep(time.Second * 2)
	log.Println("servers are running")

	coord := coordinator.NewCoordinator()
	err := coord.StartGrepping()
	if err != nil {
		log.Println("failed to process text")
		os.Exit(1)
	}
}
