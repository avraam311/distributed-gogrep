package coordinator

import "log"

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
	client := NewClient(substrs, pars)
	matches, err := client.SendAndRecieveResults()
	if err != nil {
		return err
	}
	agr := NewAgregator(matches)
	agr.PrintResultWithNumbers()
	log.Println("results are printed")

	return nil
}
