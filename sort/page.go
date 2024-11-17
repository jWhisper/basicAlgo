package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	ctx := context.Background()
	zsetKey := "user:music:favorites"

	// 初始化数据
	initData(ctx, rdb, zsetKey)

	// 初始分数为正无穷大，表示从最新的记录开始
	lastScore := "+inf"
	count := 10

	// 模拟用户请求第一页
	results, lastScore, err := getPaginatedZSetResults(ctx, rdb, zsetKey, lastScore, count)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Page 1 Results:", results)
	fmt.Println("Last Score:", lastScore)

	// 模拟 ZSET 被删除并重建
	rdb.Del(ctx, zsetKey)
	initData(ctx, rdb, zsetKey)

	// 模拟用户请求第二页
	results, lastScore, err = getPaginatedZSetResults(ctx, rdb, zsetKey, lastScore, count)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Page 2 Results:", results)
	fmt.Println("Last Score:", lastScore)
}

func initData(ctx context.Context, rdb *redis.Client, zsetKey string) {
	rdb.Del(ctx, zsetKey)
	for i := 1; i <= 100; i++ {
		member := fmt.Sprintf("music:%d", i)
		score := float64(time.Now().Unix() + int64(i))
		rdb.ZAdd(ctx, zsetKey, &redis.Z{
			Score:  score,
			Member: member,
		})
	}
}

func getPaginatedZSetResults(ctx context.Context, rdb *redis.Client, zsetKey string, lastScore string, count int) ([]string, string, error) {
	// 使用 ZREVRANGEBYSCORE 获取数据
	results, err := rdb.ZRevRangeByScore(ctx, zsetKey, &redis.ZRangeBy{
		Min:    "-inf",
		Max:    fmt.Sprintf("(%s", lastScore), // 使用“(”表示是开区间
		Count:  int64(count),
	}).Result()
	if err != nil {
		return nil, lastScore, err
	}

	// 如果没有结果，返回当前的分数
	if len(results) == 0 {
		return results, lastScore, nil
	}

	// 获取最后一条记录的分数
	lastElement := results[len(results)-1]
	score, err := rdb.ZScore(ctx, zsetKey, lastElement).Result()
	if err != nil {
		return nil, lastScore, err
	}

	// 转换分数为字符串返回
	return results, strconv.FormatFloat(score, 'f', -1, 64), nil
}
