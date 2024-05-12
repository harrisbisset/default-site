package components

import (
	"net/http"

	"github.com/harrisbisset/fnet"
)

type DefaultWrapper fnet.Component

func (a DefaultWrapper) Show(w http.ResponseWriter, req *http.Request) {
	fnet.Component(a).Render(w, req)
}
