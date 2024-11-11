package html

import (
	"github.com/lpuig/cpmanager/html/comp"
	"io"
)

func MainPage(w io.Writer) {
	mainPage := comp.Page("Ma main page")
	mainPage.Render(w)
}
