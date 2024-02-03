package cache

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

type Config struct {
	Lifetime int
}

type Cache struct {
	Client *redis.Client
	Config Config
}

func New(a ...Config) *Cache {
	c := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	conf := Config{Lifetime: 1}

	if len(a) > 0 {
		conf = a[0]
	}

	return &Cache{
		Client: c,
		Config: conf,
	}
}

func (c *Cache) Ping() (string, error) {
	ctx := context.Background()
	s, err := c.Client.Ping(ctx).Result()
	if err != nil {
		return "", err
	}

	return s, nil
}

func (c *Cache) Close() error {
	err := c.Client.Close()
	return err
}

func (c *Cache) Set(i interface{}, k string) (string, error) {
	ctx := context.Background()

	m, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	s, err := c.Client.Set(ctx, k, m, time.Duration(c.Config.Lifetime*int(time.Minute))).Result()
	if err != nil {
		return "", err
	}

	return s, nil
}

func (c *Cache) Get(i interface{}, k string) error {
	ctx := context.Background()
	r, err := c.Client.Get(ctx, k).Result()
	if err != nil {
		return err
	}

	err = json.Unmarshal([]byte(r), &i)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache) Delete(k string) (int, error) {
	ctx := context.Background()
	s, err := c.Client.Del(ctx, k).Result()
	if err != nil {
		return 0, err
	}

	return int(s), nil
}

func (c *Cache) Contains(k string) (bool, error) {
	ctx := context.Background()
	e, err := c.Client.Exists(ctx, k).Result()
	if err != nil {
		return false, err
	}

	return e == 1, nil
}
