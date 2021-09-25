package main

import (
	"bilibili_videos/download"
	"math/rand"
	"time"
)

// GetVideos 循环调用
func GetVideos(id string) {
	for {
		resList := download.GetRecommendVideos(id)

		//将时间戳设置成种子数
		rand.Seed(time.Now().Unix())
		// 从推荐列表中随机选择一个视频
		ranId := rand.Intn(len(resList))
		id = resList[ranId]

		// 别着急，小心被封ip
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	// 以 【猛男版】新 宝 岛 为起点
	// bvid := "BV1j4411W7F7"
	aid := "53851218"
	GetVideos(aid)
}
