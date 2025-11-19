package grepper

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
)

func StartServer(port string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		var flags Flags
		err := json.NewDecoder(r.Body).Decode(&flags)
		if err != nil {
			http.Error(w, fmt.Sprintf("bad request body: %v", err), http.StatusBadRequest)
			return
		}

		app := NewApp(&flags)

		results := app.Run()
		w.Header().Set("Content-Type", "application/json")
		if results == nil {
			results = [][]string{}
		}
		err = json.NewEncoder(w).Encode(results)
		if err != nil {
			http.Error(w, fmt.Sprintf("failed to encode response: %v", err), http.StatusInternalServerError)
			return
		}
	})

	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server at port %s failed: %v", port, err)
	}
}

func Run(ports [5]string) {
	var wg sync.WaitGroup

	for _, port := range ports {
		wg.Add(1)
		defer wg.Done()
		go StartServer(port)
	}

	go wg.Wait()
}
