package host

import (
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/mtg-yt/internal/site/host/components/pages"
)

func (cfg CFG) CreateRoutes() {
	s := cfg.Std.Mux

	//pages
	s.HandleComponent(fnet.GET, "/", pages.IndexHandler)
}
