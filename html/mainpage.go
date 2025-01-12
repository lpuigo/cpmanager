package html

import (
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/manager"
	"io"
)

func MainPage(w io.Writer, mgr manager.Manager) {
	mainPage := bulmacomp.Page("Ma main page", comp.ConsultantsBlock(mgr.Consultants))
	mainPage.Render(w)
}
