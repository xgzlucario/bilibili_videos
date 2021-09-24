package download

import (
	"context"
	"github.com/qiniu/qmgo"
)

var BiliColl *qmgo.Collection
var ctx = context.Background()

func init() {
	// MongoDB
	// docker: host.docker.internal
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		panic(err)
	}
	BiliColl = client.Database("bili").Collection("videos")
}
