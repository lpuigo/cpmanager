package consultant

import (
	"github.com/google/uuid"
)

type Container struct {
	index map[string]Consultant
}

func NewContainer() *Container {
	return &Container{
		index: make(map[string]Consultant),
	}

}

func (c *Container) Add(cslt Consultant) {
	id := uuid.NewString()
	cslt.Id = id
	c.index[id] = cslt
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

// Load populates this container with all persisted consultants
func (c *Container) Load() error {
	c.Add(Consultant{FistName: "John", LastName: "Doe"})
	c.Add(Consultant{FistName: "Jane", LastName: "Doe"})

	return nil
}
