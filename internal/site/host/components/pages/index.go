package pages

import (
	"github.com/harrisbisset/fnet"
	"github.com/harrisbisset/mtg-yt/internal/site/render/view/view_error"
)

var IndexHandler Page = Page{
	Component: fnet.Component{
		Name: "Index",
		Err:  view_error.DefaultError(),
	},
}
