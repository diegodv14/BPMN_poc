package controllers

import (
	"github.com/diegodv14/BPMN_poc/src/flow"
	"github.com/diegodv14/BPMN_poc/src/models"
	"github.com/gin-gonic/gin"
)

// GetWorkflow godoc
// @Summary Inicia un workflow
// @Description Inicia un workflow
// @Tags workflow
// @Produce json
// @Param name body models.Request true "Nombre del workflow"
// @Param description body models.Request true "Descripción del workflow"
// @Success 200 string message "Workflow iniciado"
// @Router /workflow [post]
func GetWorkflow(c *gin.Context) {
	var request models.Request

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "JSON inválido"})
		return
	}

	processWorkflow := flow.ProcessWorkflow{}
	processWorkflow.ProcessWorkflow(request)
	c.JSON(200, gin.H{
		"message": "Workflow iniciado",
	})
}
