package user

import (
	"encoding/json"
	"github.com/lpuig/cpmanager/persist"
	"io"
	"os"
)

type UserRecord struct {
	persist.Record
	*User
}

// NewUserRecord returns a new UserRecord
func NewUserRecord() *UserRecord {
	ur := &UserRecord{}
	ur.Record = *persist.NewRecord(func(w io.Writer) error {
		return json.NewEncoder(w).Encode(ur.User)
	})
	return ur
}

// NewRecordFrom returns a UserRecord from given user
func NewRecordFrom(u *User) *UserRecord {
	ur := NewUserRecord()
	ur.User = u
	ur.SetId(u.Login)
	return ur
}

func (ur *UserRecord) SetId(id string) {
	ur.Record.SetId(id)
	ur.User.Login = id
}

// NewUserRecordFrom returns a UserRecord populated from the given reader (Record ID is not set)
func NewUserRecordFrom(r io.Reader) (ur *UserRecord, err error) {
	ur = NewUserRecord()
	ur.User = New()
	err = json.NewDecoder(r).Decode(ur.User)
	if err != nil {
		ur = nil
		return
	}
	ur.SetId(ur.Login)
	return
}

// NewUserRecordFromFile returns a UserRecord populated and Ideed from the given file
func NewUserRecordFromFile(file string) (ur *UserRecord, err error) {
	f, err := os.Open(file)
	if err != nil {
		return
	}
	defer f.Close()
	ur, err = NewUserRecordFrom(f)
	if err != nil {
		ur = nil
		return
	}
	//id, err := ur.GetIdFromFileName(file)
	//if err != nil {
	//	ur = nil
	//	return
	//}
	//ur.SetId(id)
	ur.SetId(ur.Login)
	return
}
