package download

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

// jsoniter
var json = jsoniter.ConfigCompatibleWithStandardLibrary

// GetRecommendVideos
// 获取推荐视频列表 avid string
// 返回推荐视频av号 []string
func GetRecommendVideos(id string) []string {
	results := bson.M{}

	body, err := GetAndRead("https://api.bilibili.com/x/web-interface/view/detail?aid=" + id)
	if err != nil {
		log.Println("请求接口发生错误：", err)
		return []string{}
	}
	// 请求错误
	code := json.Get(body, "code").ToInt()
	if code < 0 {
		fmt.Println(json.Get(body, "message").ToString())
	}

	data := json.Get(body, "data")

	// 视频详细信息
	view := data.Get("View")
	results["bvid"] = view.Get("bvid").ToString()    // bv号
	results["aid"] = view.Get("aid").ToString()      // bv号
	results["tid"] = view.Get("tid").ToInt()         // 分区id
	results["tname"] = view.Get("tname").ToString()  // 分区名
	results["title"] = view.Get("title").ToString()  // 视频标题
	results["pubdate"] = view.Get("pubdate").ToInt() // 上传日期
	results["desc"] = view.Get("desc").ToString()    // 视频简介

	owner := view.Get("owner")                           // up主
	results["owner_id"] = owner.Get("mid").ToString()    // up主id
	results["owner_name"] = owner.Get("name").ToString() // up主名

	stat := view.Get("stat")                             // 视频数据
	results["view"] = stat.Get("view").ToInt64()         // 播放量
	results["danmaku"] = stat.Get("danmaku").ToInt64()   // 弹幕
	results["reply"] = stat.Get("reply").ToInt64()       // 评论
	results["favorite"] = stat.Get("favorite").ToInt64() // 收藏
	results["coin"] = stat.Get("coin").ToInt64()         // 硬币
	results["share"] = stat.Get("share").ToInt64()       // 分享
	results["his_rank"] = stat.Get("his_rank").ToInt()   // 历史全站最高排名

	fmt.Println(results["title"], "\t分区:", results["tname"], "\t作者:", results["owner_name"], "\t播放量:", fmt.Sprintf("%.1f万", float64(results["view"].(int64))/10000.0))

	_, err = BiliColl.UpsertId(ctx, results["bvid"], results)
	if err != nil {
		log.Println("写入mongo错误：", err)
	}

	//card := data.Get("Card") // UP主
	//tags := data.Get("Tags") // 标签
	//reply := data.Get("Reply") // 评论

	// 推荐视频
	related := data.Get("Related")

	aidList := make([]string, related.Size())
	for i := 0; i < related.Size(); i++ {
		// 找一个不同作者
		if related.Get(i, "tname").ToString() != "明星" {
			aidList[i] = related.Get(i, "aid").ToString()
		}
	}
	return aidList
}
