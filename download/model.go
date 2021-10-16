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

// Videos 视频信息表结构
type Videos struct {
	Bvid     string `xorm:"char(16) pk not null"` // bv号
	Pubdate  int    `xorm:"int(16) not null"`     // 上传日期
	Title    string `xorm:"text not null"`        // 视频标题
	Desc     string `xorm:"text not null"`        // 视频简介
	Duration int    `xorm:"int not null"`         // 视频时长

	View     int64 `xorm:"int not null"` // 播放量
	Danmaku  int64 `xorm:"int not null"` // 弹幕
	Like     int64 `xorm:"int not null"` // 点赞
	Reply    int64 `xorm:"int not null"` // 评论
	Favorite int64 `xorm:"int not null"` // 收藏
	Coin     int64 `xorm:"int not null"` // 投币
	Share    int64 `xorm:"int not null"` // 分享
	HisRank  int   `xorm:"int not null"` // 历史最高全站排名

	OwnerId string `xorm:"char(16) not null"` // UP主id
	Tid     int    `xorm:"int not null"`      // 分区id
}

// Uploaders UP主
type Uploaders struct {
	Id        string `xorm:"char(16) pk not null"`     // UP主id
	Name      string `xorm:"char(32) unique not null"` // UP主名
	Sex       string `xorm:"char(4) not null"`         // 性别
	Rank      int    `xorm:"int not null"`             // 全站排名
	Fans      int64  `xorm:"int not null"`             // 粉丝数
	Likes     int64  `xorm:"int not null"`             // 获赞总数
	Attention int64  `xorm:"int not null"`             // 关注
	Sign      string `xorm:"text not null"`            // 个性签名
}

// Partitions 视频分区
type Partitions struct {
	Id   int    `xorm:"int pk not null"`          // 分区id
	Name string `xorm:"char(16) unique not null"` // 分区名
}

// 建表
func init() {
	var err error

	// 连接UserDB
	connStr := "postgres://postgres:123456@bili_videos_postgres:5432/videos?sslmode=disable"
	biliDB, err = xorm.NewEngine("postgres", connStr)
	if err != nil {
		panic(err)
	}
	// 建表
	err = biliDB.Sync2(new(Videos))
	if err != nil {
		panic(err)
	}
	err = biliDB.Sync2(new(Partitions))
	if err != nil {
		panic(err)
	}
	err = biliDB.Sync2(new(Uploaders))
	if err != nil {
		panic(err)
	}

	// RedisDB
	VideosDB = redis.NewClient(&redis.Options{
		Addr: "bili_videos_redis:6379",
		DB:   0,
	})
}
