package manager

import "github.com/lpuig/cpmanager/model/consultant"

type Manager struct {
	Consultants *consultant.Container
}

func New() *Manager {
	csltCont := consultant.NewContainer()

	mgr := &Manager{
		Consultants: csltCont,
	}

	return mgr
}

func (c *Manager) Init() error {
	err := c.Consultants.Load()
	if err != nil {
		return err
	}
	return nil
}
