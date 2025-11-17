package coordinator

type Coordinator struct{}

func NewCoordinator() *Coordinator {
	return &Coordinator{}
}

func (c *Coordinator) StartGrepping() error {
	pars, err := NewParser()
	if err != nil {
		return err
	}
	sep := NewSeparator(pars.Strings)
	substrs := sep.Separate()
	client := NewClient(substrs)
	matches, err := client.SendAndRecieveResults()
	if err != nil {
		return err
	}
	agr := NewAgregator(matches)
	agr.PrintResultWithNumbers()

	return nil
}
