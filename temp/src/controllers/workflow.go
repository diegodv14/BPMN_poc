package controllers

import (
	"context"
	"log"

	"github.com/diegodv14/BPMN_poc/src/client"
	"github.com/diegodv14/BPMN_poc/src/flow"
	"github.com/diegodv14/BPMN_poc/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	sdkclient "go.temporal.io/sdk/client"
)

type WorkflowInput struct {
	InsertRequest struct {
		Returning []models.Request `json:"returning"`
	} `json:"insert_request"`
}

// GetWorkflow godoc
// @Summary Inicia un workflow
// @Description Inicia un workflow de Temporal para procesar la solicitud
// @Tags workflow
// @Accept  json
// @Produce json
// @Param request body models.Request true "Datos para iniciar el workflow"
// @Success 200 {object} map[string]string "Workflow iniciado con éxito"
// @Failure 400 {object} map[string]string "Error en la solicitud"
// @Failure 500 {object} map[string]string "Error interno del servidor"
// @Router /workflow [post]
func GetWorkflow(c *gin.Context) {
	var input WorkflowInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	if len(input.InsertRequest.Returning) == 0 {
		c.JSON(400, gin.H{"error": "El campo 'returning' está vacío o no existe"})
		return
	}

	request := input.InsertRequest.Returning[0]

	temporalClient, err := client.GetTemporalClient()
	if err != nil {
		log.Fatalln("No se pudo crear el cliente de Temporal, puede que no este corriendo el temporal", err)
		c.JSON(500, gin.H{"error": "Error interno al conectar con Temporal"})
		return
	}

	workflowOptions := sdkclient.StartWorkflowOptions{
		ID:        "process_workflow_" + uuid.New().String(),
		TaskQueue: "PROCESS_TASK_QUEUE",
	}

	wfRun, err := temporalClient.ExecuteWorkflow(context.Background(), workflowOptions, flow.ProcessWorkflow, request)
	if err != nil {
		log.Println("No se pudo iniciar el workflow", err)
		c.JSON(500, gin.H{"error": "No se pudo iniciar el workflow"})
		return
	}

	log.Println("Workflow iniciado con éxito", "WorkflowID", wfRun.GetID(), "RunID", wfRun.GetRunID())
	c.JSON(200, gin.H{
		"message":    "Workflow iniciado con éxito",
		"workflowId": wfRun.GetID(),
		"runId":      wfRun.GetRunID(),
	})
}
