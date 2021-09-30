package main

import (
	"bilibili_videos/download"
	"log"
	"sync"
	"time"
)

// 循环调用
func getVideos(id string) {
	var err error
	for {
		id, err = download.GetRecommendVideos(id)
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(time.Millisecond * 100)
	}
}

// 打印数据库信息
func printData() {
	for {
		download.ShowDataBase()
		time.Sleep(time.Minute)
	}
}

func main() {
	// 以 【猛男版】新 宝 岛 为起点
	bvid := "BV1j4411W7F7"

	group := sync.WaitGroup{}
	// 创建协程
	for gNum := 0; gNum < 3; gNum++ {
		group.Add(1)
		go func() {
			getVideos(bvid)
			group.Done()
		}()
	}
	log.Println("正在下载视频数据...")
	go printData()

	group.Wait()
}
