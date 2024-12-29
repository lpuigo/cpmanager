package manager

import (
	"github.com/lpuig/cpmanager/config"
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/model/consultant"
	"testing"
)

func TestManager_PopulateConsultant(t *testing.T) {
	conf := config.Set()
	c, err := consultant.NewConsultantPersister("../../"+conf.DirPersisterConsultant, log.New())
	if err != nil {
		t.Fatalf("NewConsultantPersister() returned unexpected error = %v", err)
	}
	c.SetPersistDelay(0)
	c.Add(&consultant.Consultant{FirstName: "John", LastName: "Doe"})
	c.Add(&consultant.Consultant{FirstName: "Jane", LastName: "Doe"})
	c.Persister.WaitPersistDone()
}
