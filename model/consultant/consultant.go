package consultant

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strings"
)

var titler cases.Caser

func init() {
	titler = cases.Title(language.French)
}

type PriceHistory struct {
	Price int
	Day   string
}

type Mission struct {
	Title     string
	Company   string
	Manager   string
	StartDay  string
	EndDay    string
	DailyCost []PriceHistory
	DailyRate []PriceHistory
}

type Consultant struct {
	Id        string
	FirstName string
	LastName  string
	Profile   string
	CrmrId    string
	Missions  []Mission
}

func NewConsultant() *Consultant {
	return &Consultant{}
}

func cleanTextAttribute(text string, title bool) string {
	res := strings.TrimSpace(text)
	if title {
		return titler.String(res)
	}
	return res
}

// Clean updates receiver with cleaned given attributes
func (c *Consultant) Clean(firstName, lastName, profile, crmid string) {
	c.FirstName = cleanTextAttribute(firstName, true)
	c.LastName = cleanTextAttribute(lastName, true)
	c.Profile = cleanTextAttribute(profile, true)
	c.CrmrId = cleanTextAttribute(crmid, false)

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

// ===  Consultant table formatting function

// Status returns receiver status, depending on its mission state (
func (c *Consultant) Status() string {
	// TODO
	return "à définir"
}

// Client returns receiver current Client, based on current date (
func (c *Consultant) Client() string {
	// TODO
	return "-"
}

// Manager returns receiver current manager, based on current date (
func (c *Consultant) Manager() string {
	// TODO
	return "-"
}

// MissionTitle returns receiver current mission title, based on current date (
func (c *Consultant) MissionTitle() string {
	// TODO
	return "-"
}

func (c *Consultant) CrmUrl() string {
	return fmt.Sprintf("https://ui.boondmanager.com/resources/%s/overview", c.CrmrId)
}
