package main

import (
	"bilibili_videos/download"
	"log"
	"time"
)

// GetVideos 循环调用
func GetVideos(id string) {
	var err error
	for {
		id, err = download.GetRecommendVideos(id)
		if err != nil {
			log.Println(err)
			return
		}
		// 别着急，小心被封ip
		time.Sleep(time.Millisecond * 50)
	}
}

func main() {
	// 以 【猛男版】新 宝 岛 为起点
	// bvid := "BV1j4411W7F7"
	aid := "53851218"
	GetVideos(aid)
}
