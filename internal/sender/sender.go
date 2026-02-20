package sender

import (
    "bytes"
    "encoding/json"
    "net/http"

    "github.com/japsieeee/agent/internal/models"
    "github.com/japsieeee/agent/internal/config"
)

func Send(metrics *models.Metrics) error {
	url := config.GetServerURL() + "/metrics"

	payload, _ := json.Marshal(metrics)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("server returned %d", resp.StatusCode)
	}

	return nil
}