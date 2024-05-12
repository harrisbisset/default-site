package site

import (
	"errors"
	"log"
	"net/http"

	"github.com/harrisbisset/mtg-yt/internal/site/host"
)

func Start() {
	cfg := host.CreateConfig()
	cfg.CreateRoutes()

	log.Printf("server is listening at %s...", cfg.Std.Port)
	err := http.ListenAndServe(cfg.Std.Port, cfg.Std.Mux.HTTPMux)

	if errors.Is(err, http.ErrServerClosed) {
		log.Print("server closed")
	} else if err != nil {
		log.Fatalf("error starting server: %s", err)
	}
}
