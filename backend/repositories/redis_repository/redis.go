package redis_repository

import (
	"context"
	"encoding/json"
	"fmt"
	"socialmedia/domain"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	redis *redis.Client
	ctx   context.Context
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{
		redis: redis,
		ctx:   context.Background(),
	}
}

func (r *RedisRepository) GetAllPost(page, limit int, user_id string) ([]domain.PostsWithUsername, error) {
	key := fmt.Sprintf("feed:page:%d:limit:%d", page, limit)
	val, err := r.redis.Get(r.ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var posts []domain.PostsWithUsername
	json.Unmarshal([]byte(val), &posts)
	return posts, nil
}

func (r *RedisRepository) CacheGetAllPost(page, limit int, user_id string, data []domain.PostsWithUsername) {
	key := fmt.Sprintf("feed:page:%d:limit:%d", page, limit)

	bytes, _ := json.Marshal(data)

	r.redis.Set(r.ctx, key, bytes, 5*time.Minute)
}

func (r *RedisRepository) DeleteFeed() error {

	val, err := r.redis.Keys(r.ctx, "feed:page:*").Result()
	if err != nil {
		return err
	}

	if len(val) > 0 {
		r.redis.Del(r.ctx, val...)
	}

	return nil
}
