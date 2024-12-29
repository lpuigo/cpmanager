package consultant

import (
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/persist"
	"sort"
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

// Get returns consultant with given ID, if found
func (cp *ConsultantsPersister) Get(id string) (*Consultant, bool) {
	cr, found := cp.GetById(id)
	if found {
		c, ok := cr.(*ConsultantRecord)
		if ok {
			return c.Consultant, true
		}
	}
	return nil, false
}

// Add adds the given ConsultantRecord to the reciever and returns the added record which id has been updated
func (cp *ConsultantsPersister) Add(ncp *Consultant) *Consultant {
	cp.Persister.Add(NewRecordFrom(ncp))
	return ncp
}

// Update updates the given ConsultantRecord
func (cp *ConsultantsPersister) Update(ncp *Consultant) error {
	return cp.Persister.Update(NewRecordFrom(ncp))

}

// Remove removes the given ConsultantRecord from the reciever
func (cp *ConsultantsPersister) Remove(ncp *Consultant) error {
	return cp.Persister.Remove(NewRecordFrom(ncp))
}

func (cp *ConsultantsPersister) GetAll() []*Consultant {
	clist := cp.GetRecords()
	res := make([]*Consultant, len(clist))
	for i, r := range clist {
		if cr, ok := r.(*ConsultantRecord); ok {
			res[i] = cr.Consultant
		}
	}
	return res
}

func (cp *ConsultantsPersister) GetSortedByName() []*Consultant {
	clist := cp.GetAll()
	sort.Slice(clist, func(i, j int) bool { return clist[i].CompareByName(clist[j]) })
	return clist
}
