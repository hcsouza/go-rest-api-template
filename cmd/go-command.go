package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hcsouza/go-rest-api-template/internal/api"
)

func main() {

	gServer := gin.New()
	api.Run(gServer)

}
