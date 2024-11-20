package manager

import "github.com/lpuig/cpmanager/model/consultant"

type Manager struct {
	consultants *consultant.Container
}

func New() *Manager {
	csltCont := consultant.NewContainer()

	mgr := &Manager{
		consultants: csltCont,
	}

	return mgr
}

func (c *Manager) Init() error {
	err := c.consultants.Load()
	if err != nil {
		return err
	}
	return nil
}
