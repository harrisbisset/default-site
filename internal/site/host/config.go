package host

import (
	"database/sql"

	"github.com/harrisbisset/fnet"
)

type (
	cfg fnet.CFG

	CFG struct {
		DB  *sql.DB
		Std cfg
	}
)

func CreateConfig() CFG {
	return CFG{
		DB:  getDatabase(loadConfig()),
		Std: cfg(fnet.CreateConfig()),
	}
}
