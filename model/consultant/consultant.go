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
