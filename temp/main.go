package main

import (
	_ "github.com/diegodv14/BPMN_poc/docs"
	"github.com/diegodv14/BPMN_poc/src/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title BPMN POC
// @version 1.0
// @description API de en Gin para probar Temporal
// @host localhost:8081
// @BasePath /

func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.WorkflowRoutes(r)
	r.Run(":8081")
}
