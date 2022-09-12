package app

import (
	"github.com/elysiumyun/elysium/internal/app/router"
	"github.com/elysiumyun/elysium/internal/pkg/datasource"
	"github.com/gin-gonic/gin"
)

func elysium() error {
	var err error

	// connect data source: like database
	ds := datasource.Get()
	ds.InitDataSource()

	// new server & init router
	engine := gin.New()
	engine.SetTrustedProxies(nil)

	router.SetRouter(engine)

	err = engine.Run(":2333")
	if err != nil {
		return err
	}
	return nil
}
