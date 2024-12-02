package manager

import (
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/model/consultant"
)

type Manager struct {
	Log         *log.Logger
	Consultants *consultant.Container
}

func New(l *log.Logger) *Manager {
	csltCont := consultant.NewContainer()

	mgr := &Manager{
		Log:         l,
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
