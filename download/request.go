package download

import (
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 5 * time.Second}

// GetAndRead 发送Get请求并读取数据 B站专用
func GetAndRead(url string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)

	headers := map[string]string{
		"cookie":     "_uuid=0299AD1E-FBDA-7575-0D88-FE80A40329B786352infoc; buvid3=36CBABEA-A71D-475A-8427-02EF770722A813414infoc; sid=4t3syq9f; buvid_fp=36CBABEA-A71D-475A-8427-02EF770722A813414infoc; DedeUserID=50868883; DedeUserID__ckMd5=657697e56f442f3d; SESSDATA=b13d8b0f%2C1636033965%2C192b4*51; bili_jct=62c68647e41d5718f180bda6e1675ea2; CURRENT_FNVAL=80; blackside_state=1; rpdid=|(kJRYmlkl|l0J'uYk|R)k|l|; LIVE_BUVID=AUTO2816205654462057; buvid_fp_plain=36CBABEA-A71D-475A-8427-02EF770722A813414infoc; CURRENT_BLACKGAP=1; fingerprint3=3de622241299b390a9ed0762b56f2a62; fingerprint=0a6ce9e9be6de4efd386d6869727ee8c; fingerprint_s=f3c6f15040c9de4a0fb2d82921723bdb; bp_video_offset_50868883=573640736909799870; CURRENT_QUALITY=120; PVID=2; bfe_id=fdfaf33a01b88dd4692ca80f00c2de7f; bp_t_offset_50868883=574130174201570980; innersign=1",
		"origin":     "https://www.bilibili.com",
		"user-agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36 Edg/93.0.961.52",
		"referer":    "https://www.bilibili.com/video/BV1j4411W7F7?from=search&seid=5221278036568599477&spm_id_from=333.337.0.0",
	}
	// 增加header选项
	for key, value := range headers {
		request.Header.Add(key, value)
	}

	res, err := myClient.Do(request)
	if err != nil {
		return nil, errors.New("请求失败！" + err.Error())
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.New("读取数据失败！" + err.Error())
	}

	return body, nil
}
