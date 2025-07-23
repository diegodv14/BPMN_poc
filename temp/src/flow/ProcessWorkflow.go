package flow

import (
	"fmt"

	"github.com/diegodv14/BPMN_poc/src/models"
)

type ProcessWorkflow struct{}

func (p *ProcessWorkflow) ProcessWorkflow(request models.Request) {
	fmt.Println(request)
}
