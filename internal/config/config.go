package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Env     string
	AgentID string
	Server  struct {
		URL           string
		Timeout       int
		Retry         struct {
			MaxAttempts    int
			BackoffSeconds int
		}
	}
	Metrics struct {
		IntervalSeconds int
		CollectCPU      bool
		CollectMemory   bool
		CollectDisk     bool
		CollectNetwork  bool
		DiskPaths       []string
	}
	Logging struct {
		Level       string
		File        string
		MaxSizeMB   int
		MaxBackups  int
		MaxAgeDays  int
	}
	Batch struct {
		Enabled           bool
		MaxBatchSize      int
		FlushIntervalSecs int
	}
}

func LoadConfig(path string) *Config {
	v := viper.New()
	v.SetConfigFile(path)
	v.AutomaticEnv() // override via ENV if needed

	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		log.Fatalf("failed to unmarshal config: %v", err)
	}

	return &cfg
}