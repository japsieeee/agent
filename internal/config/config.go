package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Env     string `mapstructure:"env"`
	AgentID string `mapstructure:"agent_id"`

	Server struct {
		URL     string `mapstructure:"url"`
		Timeout int    `mapstructure:"timeout"` // seconds

		Retry struct {
			MaxAttempts    int `mapstructure:"max_attempts"`
			BackoffSeconds int `mapstructure:"backoff_seconds"`
		} `mapstructure:"retry"`
	} `mapstructure:"server"`

	Metrics struct {
		IntervalSeconds int      `mapstructure:"interval_seconds"`
		CollectCPU      bool     `mapstructure:"collect_cpu"`
		CollectMemory   bool     `mapstructure:"collect_memory"`
		CollectDisk     bool     `mapstructure:"collect_disk"`
		CollectNetwork  bool     `mapstructure:"collect_network"`
		DiskPaths       []string `mapstructure:"disk_paths"`
	} `mapstructure:"metrics"`

	Logging struct {
		Level      string `mapstructure:"level"`
		File       string `mapstructure:"file"`
		MaxSizeMB  int    `mapstructure:"max_size_mb"`
		MaxBackups int    `mapstructure:"max_backups"`
		MaxAgeDays int    `mapstructure:"max_age_days"`
	} `mapstructure:"logging"`

	Batch struct {
		Enabled           bool `mapstructure:"enabled"`
		MaxBatchSize      int  `mapstructure:"max_batch_size"`
		FlushIntervalSecs int  `mapstructure:"flush_interval_secs"`
	} `mapstructure:"batch"`
}

func LoadConfig(path string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config: %w", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if cfg.Server.URL == "" {
		return nil, fmt.Errorf("server.url must be set")
	}

	return &cfg, nil
}

// Helper for HTTP timeout
func (c *Config) ServerTimeout() time.Duration {
	if c.Server.Timeout <= 0 {
		return 10 * time.Second
	}
	return time.Duration(c.Server.Timeout) * time.Second
}