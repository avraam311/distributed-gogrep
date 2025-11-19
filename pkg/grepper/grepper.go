package grepper

import (
	"fmt"
	"regexp"
)

type Grepper struct {
	Matches      [][]string
	Template     string
	MatchesCount int
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

	notToRepeatStringsMap := make(map[int]struct{})

	if flagV {
		for i := 0; i < length; i++ {
			if !re.Match([]byte(rows[i])) {
				g.MatchesCount++
				stringsAfter := stringsAfterLimit(flagA, length-i-1)
				stringsBefore := stringsBeforeLimit(flagB, i)
				for j := stringsBefore; j <= i+stringsAfter-1; j++ {
					if _, ok := notToRepeatStringsMap[j]; !ok {
						notToRepeatStringsMap[j] = struct{}{}
						g.Matches = append(g.Matches, []string{fmt.Sprintf("%d:", j+1), rows[j]})
					}
				}
			}
		}
		return
	}

	for i := 0; i < length; i++ {
		if re.Match([]byte(rows[i])) {
			g.MatchesCount++
			stringsAfter := stringsAfterLimit(flagA, length-i-1)
			stringsBefore := stringsBeforeLimit(flagB, i)
			for j := stringsBefore; j <= i+stringsAfter-1; j++ {
				if _, ok := notToRepeatStringsMap[j]; !ok {
					notToRepeatStringsMap[j] = struct{}{}
					g.Matches = append(g.Matches, []string{fmt.Sprintf("%d:", j+1), rows[j]})
				}
			}
		}
	}

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
		finalPosition = currentPosition
	}
	return finalPosition
}
