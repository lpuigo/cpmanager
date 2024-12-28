package consultant

import (
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/persist"
)

type ConsultantsPersister struct {
	persist.Persister
}

func NewConsultantPersister(dir string, logger *log.Logger) (*ConsultantsPersister, error) {
	cp := &ConsultantsPersister{
		Persister: *persist.NewPersister("consultant", dir, logger),
	}
	err := cp.CheckDirectory()
	if err != nil {
		return nil, err
	}
	return cp, nil
}

// LoadDirectory loads all persisted Consultants Records
func (cp *ConsultantsPersister) LoadDirectory() error {
	return cp.Persister.LoadDirectory(func(file string) (persist.Recorder, error) {
		return NewConsultantRecordFromFile(file)
	})
}

// Add adds the given ConsultantRecord to the reciever and returns the added record which id has been updated
func (cp *ConsultantsPersister) Add(ncp *ConsultantRecord) *ConsultantRecord {
	cp.Persister.Add(ncp)
	return ncp
}

// Update updates the given ConsultantRecord
//func (cp *ConsultantsPersister) Update(ncp *ConsultantRecord) error {
//	return cp.Persister.Update(ncp)
//
//}

// Remove removes the given ConsultantRecord from the reciever
//func (cp *ConsultantsPersister) Remove(ncp *ConsultantRecord) error {
//	return cp.Persister.Remove(ncp)
//}
