package routes

import (
	"github.com/diegodv14/BPMN_poc/src/controllers"
	"github.com/gin-gonic/gin"
)

func WorkflowRoutes(r *gin.Engine) {
	r.POST("/workflow", controllers.GetWorkflow)
}
