package flow

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/diegodv14/BPMN_poc/src/models"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/workflow"
)

// CallApiActivity es la actividad que realiza la llamada a la API REST.
func CallApiActivity(ctx context.Context, request models.Request) (string, error) {
	jsonData, err := json.Marshal(models.Whatsapp{
		Number: os.Getenv("WHATSAPP_NUMBER"),
		Text:   "Hola " + request.Name + " dice " + request.Description,
	})

	logger := activity.GetLogger(ctx)

	if err != nil {
		logger.Error("Error al convertir el whatsapp a JSON.", "Error", err)
		return "", err
	}

	apiURL := fmt.Sprintf("http://localhost:8084/message/sendText/%s", os.Getenv("WHATSAPP_INSTANCE"))
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error("Error al crear la solicitud HTTP.", "Error", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", "change-me")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Error("Error al realizar la llamada a Evolution API.", "Error", err)
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		logger.Error("La API devolvió un estado no esperado.", "StatusCode", resp.StatusCode)
		return "", fmt.Errorf("la API devolvió un estado no esperado: %d", resp.StatusCode)
	}

	logger.Info("Llamada a Evolution API realizada con éxito.", "StatusCode", resp.StatusCode, "Mensaje", string(jsonData))
	return "Llamada a la API exitosa", nil
}

func ProcessWorkflow(ctx workflow.Context, request models.Request) (string, error) {
	ao := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Second,
	}
	ctx = workflow.WithActivityOptions(ctx, ao)

	var result string
	err := workflow.ExecuteActivity(ctx, CallApiActivity, request).Get(ctx, &result)
	if err != nil {
		workflow.GetLogger(ctx).Error("La actividad ha fallado.", "Error", err)
		return "", err
	}

	workflow.GetLogger(ctx).Info("Workflow completado con éxito.")
	return result, nil
}
