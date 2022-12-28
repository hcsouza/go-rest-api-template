package api

import (
	"github.com/gin-gonic/gin"
	"github.com/hcsouza/go-rest-api-template/internal/api/v1/routes"
)

func Run(gServer *gin.Engine) {
	gServer.Use(
		gin.LoggerWithWriter(gin.DefaultWriter, "/health/liveness", "/health/readiness"),
		gin.Recovery(),
	)

	RegisterHealthRoutes(gServer)

	api := gServer.Group("/api")
	routes.RegisterBusinessRoutes(api)

	gServer.Run()
}
