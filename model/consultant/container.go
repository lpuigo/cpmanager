package consultant

import (
	"fmt"
	"github.com/google/uuid"
	"sort"
)

type Container struct {
	index map[string]Consultant
}

func NewContainer() *Container {
	return &Container{
		index: make(map[string]Consultant),
	}

}

func (c *Container) Add(cslt Consultant) string {
	id := uuid.NewString()
	cslt.Id = id
	c.index[id] = cslt
	return id
}

func (c *Container) Update(cslt Consultant) {
	c.index[cslt.Id] = cslt
}

func (c *Container) Get(id string) (Consultant, bool) {
	cst, ok := c.index[id]
	return cst, ok
}

func (c *Container) Remove(id string) {
	delete(c.index, id)
}

func (c *Container) GetIndex() map[string]Consultant {
	return c.index
}

func (c *Container) GetAll() []Consultant {
	clist := []Consultant{}
	for _, consultant := range c.index {
		clist = append(clist, consultant)
	}
	return clist
}

func (c *Container) GetSortedByName() []Consultant {
	clist := c.GetAll()
	sort.Slice(clist, func(i, j int) bool {
		if clist[i].LastName == clist[j].LastName {
			return clist[i].FirstName < clist[j].FirstName
		}
		return clist[i].LastName < clist[j].LastName
	})
	return clist
}

// Load populates this container with all persisted consultants
func (c *Container) Load() error {
	c.Add(Consultant{FirstName: "John", LastName: "Doe"})
	c.Add(Consultant{FirstName: "Jane", LastName: "Doe"})

	return nil
}

func (c *Container) AddNewConsultant() {
	nbC := len(c.index)
	newConsult := Consultant{
		FirstName: "John",
		LastName:  fmt.Sprintf("Doe%02d", nbC+1),
	}
	c.Add(newConsult)
}
