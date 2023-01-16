package xatu

import (
	"errors"
	"time"
)

type Config struct {
	Address            string            `yaml:"address"`
	Headers            map[string]string `yaml:"headers"`
	TLS                bool              `yaml:"tls" default:"false"`
	MaxQueueSize       int               `yaml:"max_queue_size" default:"51200"`
	BatchTimeout       time.Duration     `yaml:"batch_timeout" default:"5s"`
	ExportTimeout      time.Duration     `yaml:"export_timeout" default:"30s"`
	MaxExportBatchSize int               `yaml:"max_export_batch_size" default:"512"`
}

func (c *Config) Validate() error {
	if c.Address == "" {
		return errors.New("address is required")
	}

	return nil
}
