package grepper

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func StartServer(port string, app *App, wg *sync.WaitGroup) {
	defer wg.Done()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		app.Run()
		fmt.Fprintf(w, "Grep run at port %s\n", port)
	})
	log.Printf("Server listening at :%s\n", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Server at port %s failed: %v", port, err)
	}
}
