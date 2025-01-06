package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/zhongxic/gin-template/config"
	"github.com/zhongxic/gin-template/internal/controller/ping"
	"github.com/zhongxic/gin-template/pkg/middleware"
)

func Init(cfg *config.Config) (*gin.Engine, error) {
	return initRoutes(cfg)
}

func initRoutes(cfg *config.Config) (*gin.Engine, error) {
	r := gin.New()
	registerMiddleware(r)
	if err := registerRoutes(r, cfg); err != nil {
		return nil, err
	}
	return r, nil
}

func registerMiddleware(r *gin.Engine) {
	r.Use(middleware.Logger())
	r.Use(middleware.Recover())
}

func registerRoutes(r *gin.Engine, cfg *config.Config) error {
	pingController := &ping.Controller{}
	r.GET("/ping", pingController.Ping)
	return nil
}
