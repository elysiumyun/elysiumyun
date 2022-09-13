package app

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/elysiumyun/elysium/internal/app/router"
	"github.com/elysiumyun/elysium/internal/pkg/datasource"
	"github.com/elysiumyun/elysium/internal/pkg/middleware"
	"github.com/elysiumyun/elysium/pkg/config"
	"github.com/gin-gonic/gin"
)

func elysium() error {
	var err error
	var options config.Options = *config.Cfg().Get()

	// connect data source: like database
	ds := datasource.Get()
	ds.InitDataSource()

	// new server & init router
	engine := gin.New()
	middleware.SetMiddleware(engine)
	err = engine.SetTrustedProxies(nil)
	if err != nil {
		return nil
	}

	router.SetRouter(engine)

	var address string = strings.Join([]string{options.App.Addr, options.App.Port}, ":")

	server := &http.Server{
		Addr:           address,
		Handler:        engine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Server elysium Listening %q Success...\n", address)

	err = server.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
