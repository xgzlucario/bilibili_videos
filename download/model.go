package download

import (
	"context"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"xorm.io/xorm"
)

var (
	biliDB   *xorm.Engine
	VideosDB *redis.Client // RedisDB
	ctx      = context.Background()
)

// 视频信息表结构
type videos struct {
	Aid      string `xorm:"char(16) pk not null"`     // av号
	Bvid     string `xorm:"char(16) unique not null"` // bv号
	Tid      int    `xorm:"int not null"`             // 分区id
	Tname    string `xorm:"char(16) not null"`        // 分区名
	Pubdate  int    `xorm:"int(16) not null"`         // 上传日期
	Title    string `xorm:"text not null"`            // 视频标题
	Desc     string `xorm:"text not null"`            // 视频简介
	Duration int    `xorm:"int not null"`             // 视频时长

	View     int64 `xorm:"int not null"` // 播放量
	Danmaku  int64 `xorm:"int not null"` // 弹幕
	Like     int64 `xorm:"int not null"` // 点赞
	Reply    int64 `xorm:"int not null"` // 评论
	Favorite int64 `xorm:"int not null"` // 收藏
	Coin     int64 `xorm:"int not null"` // 投币
	Share    int64 `xorm:"int not null"` // 分享
	HisRank  int   `xorm:"int not null"` // 历史最高全站排名

	OwnerId   string `xorm:"char(16) not null"`    // UP主id
	OwnerName string `xorm:"varchar(16) not null"` // UP主名
}

func init() {
	var err error

	// 连接UserDB
	connStr := "postgres://postgres:123456@127.0.0.1:5432/videos?sslmode=disable"
	biliDB, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// 建表
	err = biliDB.Sync2(new(videos))
	if err != nil {
		panic(err)
	}

	// RedisDB
	VideosDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   5,
	})
}
