package download

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"log"
)

// jsoniter
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// GetRecommendVideos
// 获取推荐视频列表 avid string
// 返回随机视频av号 string
func GetRecommendVideos(id string) (string, error) {
	body, err := GetAndRead("https://api.bilibili.com/x/web-interface/view/detail?bvid=" + id)
	if err != nil {
		log.Println("请求接口发生错误：", err)
		return "", err
	}
	// 请求错误
	code := json.Get(body, "code").ToInt()
	if code < 0 {
		fmt.Println(json.Get(body, "message").ToString())
	}

	data := json.Get(body, "data")

	view := data.Get("View")   // 视频详细信息
	owner := view.Get("owner") // up主
	stat := view.Get("stat")   // 视频数据

	video := &videos{
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

	fmt.Println(video.Title, "\t分区:", video.Tname, "\t作者:", video.OwnerName, "\t播放量:", fmt.Sprintf("%.1f万", float64(video.View)/10000.0))

	// 先插入
	_, err = biliDB.Table("videos").Insert(video)
	if err != nil {
		// 已存在则更新
		_, err = biliDB.Table("videos").Update(video)
		if err != nil {
			log.Println("更新数据库错误：", err)
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

	// 从集合中随机读一个
	id, err = VideosDB.SRandMember(ctx, "videos").Result()
	if err != nil {
		return "", err
	}
	return id, nil
}
