package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/anggasspm/job-radar/backend/internal/dto"
	"github.com/redis/go-redis/v9"
)

// Interface 
type FavoriteCache interface {
	Get(ctx context.Context, userID uint) ([]*dto.FavoriteDetail, error)
	Set(ctx context.Context, userID uint, favorites []*dto.FavoriteDetail) error
	Delete(ctx context.Context, userID uint) error
}

type favoriteCache struct {
	client *redis.Client
	ttl    time.Duration
}

// Injection
func NewFavoriteCache(client *redis.Client) FavoriteCache {
	return &favoriteCache{
		client: client,
		ttl:    5 * time.Minute,
	}
}

func favoriteKey(userID uint) string {
	return fmt.Sprintf("favorite:%d", userID)
}

// Get data function, using context to pass data life-cycle requests
func (c *favoriteCache) Get(ctx context.Context, userID uint) ([]*dto.FavoriteDetail, error) {
	result, err := c.client.Get(ctx, favoriteKey(userID)).Result()
	if err != nil {
		return nil, err
	}

	var favorites []*dto.FavoriteDetail

	// json -> struct go
	err = json.Unmarshal([]byte(result), &favorites)
	if err != nil {
		return nil, err
	}

	return favorites, nil
}

// Set data to Redis
func (c *favoriteCache) Set(
	ctx context.Context,
	userID uint,
	favorites []*dto.FavoriteDetail,
) error {

	// struct go -> json
	data, err := json.Marshal(favorites)
	if err != nil {
		return err
	}

	return c.client.Set(
		ctx,
		favoriteKey(userID),
		data,
		c.ttl,
	).Err()
}

// delete data from Redis
func (c *favoriteCache) Delete(ctx context.Context, userID uint) error {
	return c.client.Del(ctx, favoriteKey(userID)).Err()
}