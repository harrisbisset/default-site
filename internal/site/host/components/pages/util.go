package pages

import (
	"net/http"

	"github.com/harrisbisset/fnet"
)

type Page struct {
	Component fnet.Component
}

func (p Page) Show(w http.ResponseWriter, req *http.Request) {
	fnet.Component(p.Component).Render(w, req)
}
