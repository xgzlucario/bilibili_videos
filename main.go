package main

import (
	"bilibili_videos/download"
	"fmt"
	"log"
	"sync"
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
		time.Sleep(time.Millisecond * 100)
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
			GetVideos(bvid)
			group.Done()
		}()
	}
	fmt.Println("程序正在运行中...")
	group.Wait()
}
