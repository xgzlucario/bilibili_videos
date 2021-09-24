package main

import (
	"bilibili_videos/download"
	"fmt"
)

// GetVideos 循环调用
// 这里不使用多协程，防止被封ip
func GetVideos(id string) {
	idList := download.GetRecommendVideos(id)
	for _, id = range idList {
		// 总长度
		fmt.Println(len(idList))

		resList := download.GetRecommendVideos(id)
		for _, i := range resList {
			idList = append(idList, i)
		}
	}
}

func main() {
	// 以 【猛男版】新 宝 岛 为起点
	// bvid := "BV1j4411W7F7"
	aid := "53851218"
	GetVideos(aid)
}
