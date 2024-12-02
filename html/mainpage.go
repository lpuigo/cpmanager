package html

import (
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/manager"
	"io"
)

func MainPage(w io.Writer, mgr manager.Manager) {
	mainPage := comp.Page("Ma main page", comp.ConsultantsTable(mgr.Consultants))
	mainPage.Render(w)
}
