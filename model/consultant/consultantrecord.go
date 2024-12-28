package consultant

import (
	"encoding/json"
	"github.com/lpuig/cpmanager/persist"
	"io"
	"os"
)

type ConsultantRecord struct {
	persist.Record
	*Consultant
}

// NewConsultantRecord returns a new ConsultantRecord
func NewConsultantRecord() *ConsultantRecord {
	cr := &ConsultantRecord{}
	cr.Record = *persist.NewRecord(func(w io.Writer) error {
		return json.NewEncoder(w).Encode(cr.Consultant)
	})
	return cr
}

// NewRecordFrom returns a ConsultantRecord from given consultant
func NewRecordFrom(c *Consultant) *ConsultantRecord {
	cr := NewConsultantRecord()
	cr.Consultant = c
	cr.SetId(c.Id)
	return cr
}

func (cr *ConsultantRecord) SetId(id string) {
	cr.Record.SetId(id)
	cr.Consultant.Id = id
}

// NewConsultantRecordFrom returns a ConsultantRecord populated from the given reader (Record ID is not set)
func NewConsultantRecordFrom(r io.Reader) (cr *ConsultantRecord, err error) {
	cr = NewConsultantRecord()
	err = json.NewDecoder(r).Decode(cr)
	if err != nil {
		cr = nil
		return
	}
	cr.SetId(cr.Id)
	return
}

// NewConsultantRecordFromFile returns a ConsultantRecord populated and Ideed from the given file
func NewConsultantRecordFromFile(file string) (cr *ConsultantRecord, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	cr, err = NewConsultantRecordFrom(f)
	if err != nil {
		cr = nil
		return
	}
	id, err := cr.GetIdFromFileName(file)
	if err != nil {
		cr = nil
		return
	}
	cr.SetId(id)
	return
}
