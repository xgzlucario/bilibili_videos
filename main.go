package main

import (
	"bilibili_videos/download"
	"log"
	"sync"
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
	}
}

func main() {
	// 以 【猛男版】新 宝 岛 为起点
	bvid := "BV1j4411W7F7"

	group := sync.WaitGroup{}
	// 创建协程
	for gNum := 0; gNum < 2; gNum++ {

		group.Add(1)
		go func() {
			GetVideos(bvid)
			group.Done()
		}()
	}
	group.Wait()
}
