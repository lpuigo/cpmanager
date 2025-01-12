package consultant

import "github.com/lpuig/cpmanager/model/date"

type PriceHistory struct {
	Price int
	Day   string
}

func NewPriceHistory(price int, day string) PriceHistory {
	return PriceHistory{
		Price: price,
		Day:   day,
	}
}

func GetLast(phs []PriceHistory) *PriceHistory {
	return &phs[len(phs)-1]
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

func NewMission() *Mission {
	return &Mission{
		DailyCost: []PriceHistory{{
			Price: 0,
			Day:   date.Today().String(),
		}},
		DailyRate: []PriceHistory{{
			Price: 0,
			Day:   date.Today().String(),
		}},
	}
}

func (m Mission) IsActive() bool {
	today := date.Now().String()
	if today < m.StartDay {
		return false
	}
	if m.EndDay == "" {
		return true
	}
	if today > m.EndDay {
		return false
	}
	return true
}

func NewCleanMission(title, company, manager, startDay, endDay, dailyCost, dailyRate string) Mission {
	csd := cleanTextAttribute(startDay, false)
	cp := cleanIntAttribute(dailyCost)
	rp := cleanIntAttribute(dailyRate)

	return Mission{
		Title:     cleanTextAttribute(title, true),
		Company:   cleanTextAttribute(company, true),
		Manager:   cleanTextAttribute(manager, true),
		StartDay:  csd,
		EndDay:    cleanTextAttribute(endDay, false),
		DailyCost: []PriceHistory{NewPriceHistory(cp, csd)},
		DailyRate: []PriceHistory{NewPriceHistory(rp, csd)},
	}

}
