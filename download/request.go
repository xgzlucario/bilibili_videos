package download

import (
	"errors"
	"io/ioutil"
	"net/http"
)

// GetAndRead 发送Get请求并读取数据
func GetAndRead(url string) ([]byte, error) {
	res, err := http.Get(url)

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
