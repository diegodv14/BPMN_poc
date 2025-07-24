package main

import (
	"log"

	"github.com/diegodv14/BPMN_poc/src/client"
	"github.com/diegodv14/BPMN_poc/src/flow"
	"github.com/joho/godotenv"
	"go.temporal.io/sdk/worker"

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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c, err := client.GetTemporalClient()
	if err != nil {
		log.Fatalln("Unable to create client", err)
	}

	w := worker.New(c, "PROCESS_TASK_QUEUE", worker.Options{})
	w.RegisterWorkflow(flow.ProcessWorkflow)
	w.RegisterActivity(flow.CallApiActivity)

	go func() {
		err = w.Run(worker.InterruptCh())
		if err != nil {
			log.Fatalln("Unable to start worker", err)
		}
	}()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	routes.WorkflowRoutes(r)
	r.Run(":8081")
}
