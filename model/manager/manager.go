package manager

import (
	"fmt"
	"github.com/lpuig/cpmanager/config"
	"github.com/lpuig/cpmanager/http/session"
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/model/consultant"
	"github.com/lpuig/cpmanager/model/user"
)

type Manager struct {
	Log      log.Logger
	User     user.User
	LoggedIn bool

	Sessions    *session.Sessions
	Consultants *consultant.ConsultantsPersister
}

func New(l *log.Logger, conf config.Config) (*Manager, error) {
	csltCont, err := consultant.NewConsultantPersister(conf.DirPersisterConsultant, l)
	if err != nil {
		return nil, fmt.Errorf("could not create consultant persister: %s", err.Error())
	}
	mgr := &Manager{
		Log:      *l,
		User:     *user.New(),
		LoggedIn: false,

		Sessions:    session.NewWithConfig(conf),
		Consultants: csltCont,
	}

	return mgr, nil
}

// Init populates all manager containers
func (c *Manager) Init() error {
	err := c.Consultants.LoadDirectory()
	if err != nil {
		return err
	}
	return nil
}

// Clone returns a clone the reciever, with local Log and empty User attributes
func (m Manager) Clone() Manager {
	// Log and User are local instances
	m.User = *user.New()
	m.LoggedIn = false
	// others are shared pointers (sessions and persisters)
	return m
}
