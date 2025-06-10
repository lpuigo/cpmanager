package html

import (
	"github.com/lpuig/cpmanager/html/bulmacomp"
	"github.com/lpuig/cpmanager/html/comp"
	"github.com/lpuig/cpmanager/model/manager"
	"io"
)

func MainPage(mgr manager.Manager, w io.Writer) {
	mainPage := bulmacomp.Page("Ma main page", mgr, comp.ConsultantsBlock(mgr.Consultants))
	mainPage.Render(w)
}
