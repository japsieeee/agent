package sender

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/japsieeee/agent/internal/config"
	"github.com/japsieeee/agent/internal/models"
)

type Sender struct {
	client *http.Client
	cfg    *config.Config
}

func NewSender(cfg *config.Config) *Sender {
	return &Sender{
		cfg: cfg,
		client: &http.Client{
			Timeout: cfg.ServerTimeout(),
		},
	}
}

func (s *Sender) Send(metrics *models.Metrics) error {
	if metrics == nil {
		return fmt.Errorf("metrics cannot be nil")
	}

	payload, err := json.Marshal(metrics)
	if err != nil {
		return fmt.Errorf("failed to marshal metrics: %w", err)
	}

	url := strings.TrimRight(s.cfg.Server.URL, "/") + "/metrics"

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("server returned status %d", resp.StatusCode)
	}

	return nil
}