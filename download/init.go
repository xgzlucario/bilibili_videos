package download

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/qiniu/qmgo"
)

var BiliColl *qmgo.Collection // MongoDB
var VideosDB *redis.Client    // RedisDB
var ctx = context.Background()

func init() {
	// MongoDB
	// docker: host.docker.internal
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		panic(err)
	}
	BiliColl = client.Database("bili").Collection("videos")

	// RedisDB
	VideosDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   5,
	})
}
