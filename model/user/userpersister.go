package user

import (
	"github.com/lpuig/cpmanager/log"
	"github.com/lpuig/cpmanager/persist"
)

type UsersPersister struct {
	persist.Persister
}

func NewUserPersister(dir string, logger *log.Logger) (*UsersPersister, error) {
	up := &UsersPersister{
		Persister: *persist.NewPersister("user", dir, logger),
	}
	err := up.CheckDirectory()
	if err != nil {
		return nil, err
	}
	return up, nil
}

// LoadDirectory loads all persisted Users Records
func (up *UsersPersister) LoadDirectory() error {
	return up.Persister.LoadDirectory(func(file string) (persist.Recorder, error) {
		return NewUserRecordFromFile(file)
	})
}

// Get returns user with given ID, if found
func (up *UsersPersister) Get(id string) (*User, bool) {
	ur, found := up.GetById(id)
	if found {
		u, ok := ur.(*UserRecord)
		if ok {
			return u.User, true
		}
	}
	return nil, false
}

// Add adds the given UserRecord to the reciever and returns the added record which id has been updated
func (up *UsersPersister) Add(nup *User) *User {
	//TODO ensure user id uniqueness
	up.Persister.Add(NewRecordFrom(nup))
	return nup
}

// Update updates the given UserRecord
func (up *UsersPersister) Update(nup *User) error {
	return up.Persister.Update(NewRecordFrom(nup))
}

// Remove removes the given UserRecord from the reciever
func (up *UsersPersister) Remove(nup *User) error {
	return up.Persister.Remove(NewRecordFrom(nup))
}

func (up *UsersPersister) GetAll() []*User {
	ulist := up.GetRecords()
	res := make([]*User, len(ulist))
	for i, r := range ulist {
		if ur, ok := r.(*UserRecord); ok {
			res[i] = ur.User
		}
	}
	return res
}

// ValidateCredentials validates the user credentials
func (up *UsersPersister) ValidateCredentials(username, password string) (*User, bool) {
	// First check if the user exists in the persister
	user, found := up.Get(username)
	if found && user.Password == password {
		return user, true
	}

	// TODO Remove fallback
	// Fallback to hardcoded admin user if no users are found or credentials don't match
	if username == "admin" && password == "password" {
		return &User{
			FullName: "Administrator",
			Login:    "admin",
			Password: "password", // In a real application, you would never store passwords in plain text
		}, true
	}
	return nil, false
}
