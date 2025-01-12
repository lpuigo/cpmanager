package consultant

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"sort"
	"strconv"
	"strings"
)

var titler cases.Caser

func init() {
	titler = cases.Title(language.French)
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

func cleanIntAttribute(text string) int {
	ttxt := strings.TrimSpace(text)
	res, err := strconv.Atoi(ttxt)
	if err != nil {
		return 0
	}
	return res
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

func (c *Consultant) HasActiveMission() bool {
	return len(c.Missions) > 0 && c.LastMission().IsActive()
}

func (c *Consultant) LastMission() *Mission {
	if len(c.Missions) == 0 {
		return nil
	}
	return &c.Missions[len(c.Missions)-1]
}

func (c *Consultant) AddMission(m Mission) {
	c.Missions = append(c.Missions, m)
	c.sortMission()
}

func (c *Consultant) sortMission() {
	sort.Slice(c.Missions, func(i, j int) bool {
		return c.Missions[i].StartDay < c.Missions[j].StartDay
	})
}
