package consultant

type Consultant struct {
	Id        string
	FirstName string
	LastName  string
}

// Name returns normed consultant's name (last- first-name)
func (c Consultant) Name() string {
	return c.FirstName + " " + c.LastName
}

func (c *Consultant) CompareByName(c2 *Consultant) bool {
	if c.LastName == c2.LastName {
		return c.FirstName < c2.FirstName
	}
	return c.LastName < c2.LastName
}
