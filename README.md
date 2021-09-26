# bilibili_videos

## 一、项目介绍

本项目为爬取b站所有视频数据，用于个人统计

同时还是一个测试docker相关内容的项目，包括Dockerfile、容器互联、docker-compose、数据卷、网络相关知识

api地址：https://api.bilibili.com/x/web-interface/view/detail?bvid=BV1j4411W7F7

#### 数据结构

```go
// Videos 视频信息表结构
type Videos struct {
	Bvid     string `xorm:"char(16) pk not null"` // bv号
	Tid      int    `xorm:"int not null"`         // 分区id
	Tname    string `xorm:"char(16) not null"`    // 分区名
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

	OwnerId   string `xorm:"char(16) not null"` // UP主id
	OwnerName string `xorm:"char(32) not null"` // UP主名
}
```

#### 数据库存储路径

数据库信息使用docker数据卷！

所有视频bv号：docker-volume-redis_data

视频详细信息：docker-volume-pg_data

