package coordinator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

type Client struct {
	substrs [][]string
	pars    *Parser
}

func NewClient(substrs [][]string, pars *Parser) *Client {
	return &Client{
		substrs: substrs,
		pars:    pars,
	}
}

func (c *Client) SendAndRecieveResults() ([][]string, error) {
	type result struct {
		data [][]string
		err  error
	}

	resultsChan := make(chan result, len(c.substrs))
	quorum := len(c.substrs)/2 + 1

	serverAddrs := []string{
		"http://localhost:8080",
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
		"http://localhost:8084",
	}

	var wg sync.WaitGroup

	for i, substr := range c.substrs {
		pars := Parser{
			FlagA:    c.pars.FlagA,
			FlagB:    c.pars.FlagB,
			FlagC:    c.pars.FlagC,
			FlagCc:   c.pars.FlagCc,
			FlagI:    c.pars.FlagI,
			FlagV:    c.pars.FlagV,
			FlagF:    c.pars.FlagF,
			FlagN:    c.pars.FlagN,
			Template: c.pars.Template,
			Strings:  substr,
		}

		wg.Add(1)
		go func(addr string, data *Parser) {
			defer wg.Done()

			jsonData, err := json.Marshal(data)
			if err != nil {
				resultsChan <- result{nil, fmt.Errorf("json marshal: %w", err)}
				return
			}

			req, err := http.NewRequest("POST", addr, bytes.NewBuffer(jsonData))
			if err != nil {
				resultsChan <- result{nil, fmt.Errorf("new request: %w", err)}
				return
			}
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				resultsChan <- result{nil, fmt.Errorf("http do: %w", err)}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				resultsChan <- result{nil, fmt.Errorf("bad status: %s", resp.Status)}
				return
			}

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				resultsChan <- result{nil, fmt.Errorf("read body: %w", err)}
				return
			}

			var serverResult [][]string
			if err := json.Unmarshal(body, &serverResult); err != nil {
				resultsChan <- result{nil, fmt.Errorf("unmarshal: %w", err)}
				return
			}

			resultsChan <- result{serverResult, nil}

		}(serverAddrs[i], &pars)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	var collected [][]string
	successCount := 0

	for res := range resultsChan {
		if res.err == nil {
			collected = append(collected, res.data...)
			successCount++
			if successCount >= quorum {
				break
			}
		} else {
			log.Printf("error from server: %v", res.err)
		}
	}

	if successCount < quorum {
		return nil, fmt.Errorf("failed to reach quorum: only %d/%d successful", successCount, len(c.substrs))
	}

	return collected, nil
}
