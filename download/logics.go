package download

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
	"time"
)

// jsoniter
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// 获取随机视频bvid
func getRandomBvid() string {
	// 从集合中随机读一个
	id, _ := VideosDB.SRandMember(ctx, "videos").Result()
	return id
}

// GetRecommendVideos
// id:    视频bv号
// return: 随机视频bv号
func GetRecommendVideos(id string) (string, error) {
GET:
	body, err := GetAndRead("http://api.bilibili.com/x/web-interface/view/detail?bvid=" + id)
	if err != nil {
		log.Println(err, "3秒后重试...")
		time.Sleep(time.Second * 3)
		goto GET
	}
	// 请求错误
	code := json.Get(body, "code").ToInt()
	if code < 0 {
		// 打印错误信息, 更换id
		log.Println(json.Get(body, "message").ToString())
		log.Println(err, "3秒后重试...")
		time.Sleep(time.Second * 3)
		return getRandomBvid(), nil
	}

	data := json.Get(body, "data")

	view := data.Get("View")   // 视频详细信息
	owner := view.Get("owner") // up主
	stat := view.Get("stat")   // 视频数据

	video := &Videos{
		Bvid:     view.Get("bvid").ToString(),  // bv号
		Tid:      view.Get("tid").ToInt(),      // 分区id
		Tname:    view.Get("tname").ToString(), // 分区名
		Title:    view.Get("title").ToString(), // 视频标题
		Pubdate:  view.Get("pubdate").ToInt(),  // 上传日期
		Desc:     view.Get("desc").ToString(),  // 视频简介
		Duration: view.Get("duration").ToInt(), // 视频时长

		View:     stat.Get("view").ToInt64(),     // 播放量
		Like:     stat.Get("like").ToInt64(),     // 点赞
		Danmaku:  stat.Get("danmaku").ToInt64(),  // 弹幕
		Reply:    stat.Get("reply").ToInt64(),    // 评论
		Favorite: stat.Get("favorite").ToInt64(), // 收藏
		Coin:     stat.Get("coin").ToInt64(),     // 硬币
		Share:    stat.Get("share").ToInt64(),    // 分享
		HisRank:  stat.Get("his_rank").ToInt(),   // 历史全站最高排名

		OwnerId:   owner.Get("mid").ToString(),  // up主id
		OwnerName: owner.Get("name").ToString(), // up主名
	}

	// 先插入
	_, err = biliDB.Table("videos").Insert(video)
	if err != nil {
		// 已存在则更新
		_, err = biliDB.Table("videos").ID(video.Bvid).Update(video)
		if err != nil {
			log.Println("update db error: ", video, err)
		}
	}

	//card := data.Get("Card") // UP主
	//tags := data.Get("Tags") // 标签
	//reply := data.Get("Reply") // 评论

	// 推荐视频
	related := data.Get("Related")

	for i := 0; i < related.Size(); i++ {
		// 添加至RedisDB 集合
		VideosDB.SAdd(ctx, "videos", related.Get(i, "bvid").ToString())
	}

	// 获取随机id
	return getRandomBvid(), nil
}

// ShowDataBase 展示数据库信息
func ShowDataBase() {
	count1, _ := biliDB.Table("videos").Count()
	count2, _ := VideosDB.SCard(ctx, "videos").Result()
	fmt.Println("PostgreSql 数据总量:", count1)
	fmt.Println("Redis 数据总量:", count2)
}
