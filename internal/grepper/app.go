package grepper

import (
	"fmt"

	"github.com/avraam311/distributed-gogrep/internal/coordinator"
)

type App struct {
	grepper *Grepper
	flags   *coordinator.Parser
}

func NewApp(grep *Grepper, flags *coordinator.Parser) *App {
	return &App{
		grepper: grep,
		flags:   flags,
	}
}

func (a *App) Run() {
	if a.flags.FlagF {
		a.grepper.ConsiderTemplateAsString()
	}

	if a.flags.FlagI {
		a.grepper.IgnoreCase()
	}

	a.grepper.Grep(a.flags.FlagA, a.flags.FlagB, a.flags.FlagC, a.flags.FlagV)

	if a.flags.FlagCc {
		fmt.Println(a.grepper.MatchesCount)
		return
	}

	if a.flags.FlagN {
		a.grepper.PrintResultWithNumbers()
		return
	}

	a.grepper.PrintResult()
}
