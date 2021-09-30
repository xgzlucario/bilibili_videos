# bilibili_videos

### 项目介绍

本项目为爬取b站所有视频数据，用于个人统计

同时还是一个测试docker相关内容的项目，包括Dockerfile、容器互联、docker-compose、数据卷、网络相关知识

api地址：https://api.bilibili.com/x/web-interface/view/detail?bvid=BV1j4411W7F7



### 启动项目

打开cmd，进入项目根目录，输入

```shell
docker-compose up
```

得到日志

```shell
2021/09/30 14:20:53 正在下载视频数据...
2021/09/30 14:20:53 Postgresql数据总量: 0
2021/09/30 14:20:53 Redis数据总量: 0
2021/09/30 14:20:56 访问权限不足
2021/09/30 14:20:56 <nil> 3秒后重试...
2021/09/30 14:21:53 Postgresql数据总量: 750
2021/09/30 14:21:53 Redis数据总量: 14554
2021/09/30 14:22:53 Postgresql数据总量: 1495
2021/09/30 14:22:53 Redis数据总量: 28443
2021/09/30 14:23:53 Postgresql数据总量: 2215
2021/09/30 14:23:53 Redis数据总量: 42210
2021/09/30 14:24:53 Postgresql数据总量: 2940
2021/09/30 14:24:53 Redis数据总量: 55863
...
```



### 数据库存储路径

所有视频bv号：.data/redis

视频详细信息：.data/pg



### 如何查看所有视频数据

首先启动容器，容器中的postgres数据库自动加载数据，并映射到本地端口10123

```shell
# docker-compose.yml
ports:
  	- "10123:5432"  # 暴露端口 可以在本地端口10123查看数据
volumes:
	- .data/pg:/var/lib/postgresql  # 数据持久化至本地
environment:
	- POSTGRES_DB=videos  # 默认数据库
	- POSTGRES_USER=postgres  # 用户名
	- POSTGRES_PASSWORD=123456  # 密码
```

打开navicat或其他数据库管理工具，输入端口、数据库名（"videos"）、用户名、密码，即可连接



### 数据结构

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

