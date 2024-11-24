package consultant

type Consultant struct {
	Id       string
	FistName string
	LastName string
}

func (c Consultant) GetId() string {
	return c.Id
}

// Name returns normed consultant's name (last- first-name)
func (c Consultant) Name() string {
	return c.FistName + " " + c.LastName
}
