package cluster

import (
	"context"
	"errors"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

const Type = "redis-cluster"

type Cluster struct {
	config *Config

	log logrus.FieldLogger

	options *redis.ClusterOptions
	client  *redis.ClusterClient

	metrics *Metrics
}

func New(config *Config, log logrus.FieldLogger) (*Cluster, error) {
	options, err := redis.ParseClusterURL(config.Address)
	if err != nil {
		return nil, err
	}

	return &Cluster{
		config:  config,
		log:     log.WithField("store/cache", Type),
		options: options,
		metrics: NewMetrics("xatu_server_store_cache"),
	}, nil
}

func (c *Cluster) Type() string {
	return Type
}

func (c *Cluster) Start(ctx context.Context) error {
	c.client = redis.NewClusterClient(c.options)

	return nil
}

func (c *Cluster) Stop(ctx context.Context) error {
	if c.client != nil {
		if err := c.client.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Cluster) Get(ctx context.Context, key string) (*string, error) {
	cmd := c.client.Get(ctx, key)
	if cmd.Err() != nil {
		if errors.Is(cmd.Err(), redis.Nil) {
			c.metrics.AddGet(1, c.Type(), "miss")
			return nil, nil
		}

		c.metrics.AddGet(1, c.Type(), "error")

		return nil, cmd.Err()
	}

	c.metrics.AddGet(1, c.Type(), "hit")

	item := cmd.Val()

	return &item, nil
}

func (c *Cluster) Set(ctx context.Context, key, value string, ttl time.Duration) error {
	cmd := c.client.Set(ctx, key, value, ttl)

	if cmd.Err() != nil {
		c.metrics.AddSet(1, c.Type(), "error")
		return cmd.Err()
	}

	c.metrics.AddSet(1, c.Type(), "ok")

	return nil
}

func (c *Cluster) Delete(ctx context.Context, key string) error {
	cmd := c.client.Del(ctx, key)

	if cmd.Err() != nil {
		c.metrics.AddDelete(1, c.Type(), "error")
		return cmd.Err()
	}

	c.metrics.AddDelete(1, c.Type(), "ok")

	return nil
}