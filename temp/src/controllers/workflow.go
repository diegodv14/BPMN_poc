package controllers

import "github.com/gin-gonic/gin"

// GetWorkflow godoc
// @Summary Inicia un workflow
// @Description Inicia un workflow
// @Tags workflow
// @Produce json
// @Success 200 string message "Workflow iniciado"
// @Router /workflow [post]
func GetWorkflow(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Workflow iniciado",
	})
}
