package grepper

type Flags struct {
	FlagA    int
	FlagB    int
	FlagC    int
	FlagCc   bool
	FlagI    bool
	FlagV    bool
	FlagF    bool
	FlagN    bool
	Template string
	Strings  []string
}

type App struct {
	*Grepper
	flags *Flags
}

func NewApp(flags *Flags) *App {
	g := NewGrepper(flags.Strings, flags.Template)
	return &App{
		Grepper: g,
		flags:   flags,
	}
}

func (a *App) Run() [][]string {
	if a.flags.FlagF {
		a.ConsiderTemplateAsString()
	}

	if a.flags.FlagI {
		a.IgnoreCase()
	}

	a.Grep(a.flags.FlagA, a.flags.FlagB, a.flags.FlagC, a.flags.FlagV, a.flags.Strings)

	if a.flags.FlagCc {
		return nil
	}

	return a.Matches
}
