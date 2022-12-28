package main

import (
	"go-web-demo/src/configs"
	"go-web-demo/src/db"
	"go-web-demo/src/handlers"
	"go-web-demo/src/loggers"
	"go-web-demo/src/models"
	"go-web-demo/src/routers"
	"go-web-demo/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	var err error

	conf, err := configs.InitConfig(configs.GetConfigEnv())
	if err != nil {
		panic(err)
	}

	sugaredLogger, logger := loggers.InitLogger(conf.LogConfig)

	db, err := db.GormInit(conf.DBConfig, db.TableSlice, sugaredLogger)
	if err != nil {
		panic(err)
	}

	server, err := services.NewServer(services.WithConfig(conf),
		services.WithGinEngin(),
		services.WithGormDb(db),
		services.WithLog(logger),
		services.WithSuLog(sugaredLogger),
	)
	if err != nil {
		panic(err)
	}

	err = Start(server)
	if err != nil {
		panic(err)
	}

}

func Start(s *services.Server) error {
	//loading middleware
	s.GetGinEngine().Use(handlers.Cors())
	s.GetGinEngine().Use(loggers.GinLogger(s.GetLogger()))
	s.GetGinEngine().Use(loggers.GinRecovery(s.GetLogger(), true))
	s.GetGinEngine().GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg":  models.RESP_CODE_SUCCESS,
			"data": "Hello,World!",
		})
	})

	//loading route

	routers.LoadNoTokenRouter(s)

	routers.LoadMonitorRouter(s)

	s.GetGinEngine().Use(handlers.JWTAuthMiddleware(s))

	routers.LoadUserRouter(s)

	routers.LoadControlRouter(s)

	routers.LoadExecRouter(s)

	routers.LoadTraceRouter(s)

	err := s.GetGinEngine().Run(":" + s.GetSeverPort())
	if err != nil {
		return err
	}

	return nil
}
