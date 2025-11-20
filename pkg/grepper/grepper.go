package grepper

import (
	"fmt"
	"regexp"
	"sync"
)

type Grepper struct {
	Matches      [][]string
	Template     string
	MatchesCount int
	mu           sync.Mutex
}

func NewGrepper(rows []string, template string) *Grepper {
	return &Grepper{
		Matches:      [][]string{},
		Template:     template,
		MatchesCount: 0,
	}
}

func (g *Grepper) Grep(flagA int, flagB int, flagC int, flagV bool, rows []string) {
	re := regexp.MustCompile(g.Template)
	if flagC != 0 && (flagA == 0 && flagB == 0) {
		flagA = flagC
		flagB = flagC
	}
	length := len(rows)

	notToRepeatStringsMap := struct {
		m map[int]struct{}
		sync.Mutex
	}{m: make(map[int]struct{})}

	var wg sync.WaitGroup

	worker := func(start, end int) {
		defer wg.Done()
		for i := start; i < end; i++ {
			matchFound := re.Match([]byte(rows[i]))
			if flagV {
				matchFound = !matchFound
			}
			if matchFound {
				g.mu.Lock()
				g.MatchesCount++
				g.mu.Unlock()

				stringsAfter := stringsAfterLimit(flagA, length-i-1)
				stringsBefore := stringsBeforeLimit(flagB, i)

				for j := stringsBefore; j <= i+stringsAfter-1; j++ {
					notToRepeatStringsMap.Lock()
					if _, ok := notToRepeatStringsMap.m[j]; !ok {
						notToRepeatStringsMap.m[j] = struct{}{}
						notToRepeatStringsMap.Unlock()
						g.mu.Lock()
						g.Matches = append(g.Matches, []string{fmt.Sprintf("%d:", j+1), rows[j]})
						g.mu.Unlock()
					} else {
						notToRepeatStringsMap.Unlock()
					}
				}
			}
		}
	}

	numWorkers := 8
	chunkSize := (length + numWorkers - 1) / numWorkers

	for w := 0; w < numWorkers; w++ {
		start := w * chunkSize
		end := start + chunkSize
		if end > length {
			end = length
		}
		if start >= length {
			break
		}
		wg.Add(1)
		go worker(start, end)
	}

	wg.Wait()
}

func (g *Grepper) ConsiderTemplateAsString() {
	g.Template = regexp.QuoteMeta(g.Template)
}

func (g *Grepper) IgnoreCase() {
	g.Template = `(?i)` + g.Template
}

func stringsAfterLimit(stringsAfter int, stringsLeft int) int {
	if stringsLeft < stringsAfter {
		stringsAfter = stringsLeft
	}
	return stringsAfter + 1
}

func stringsBeforeLimit(stringsBefore int, currentPosition int) int {
	finalPosition := currentPosition - stringsBefore
	if finalPosition < 0 {
		finalPosition = 0
	}
	return finalPosition
}
