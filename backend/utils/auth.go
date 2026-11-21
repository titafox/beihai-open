package utils

import (
	"os"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type Item struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		State string `json:"state"`
	} `json:"data"`
}

// 实名认证
func CheckIdCard(idcard, name, mobile string) (bool, error) {
	const path = "http://sjsys.market.alicloudapi.com/communication/personal/1979"
	values := url.Values{}
	values.Add("idcard", idcard)
	values.Add("name", name)
	values.Add("mobile", mobile)
	reader := strings.NewReader(values.Encode())
	req, _ := http.NewRequest(http.MethodPost, path, reader)
	req.Header.Add("Authorization", "APPCODE "+os.Getenv("ALIYUN_MARKET_APPCODE"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, errors.New("解析实名结果错误")
	}

	var item Item
	err = json.Unmarshal(body, &item)
	if err != nil {
		return false, errors.New("实名遇到问题:" + err.Error())
	}
	fmt.Println(item)

	if item.Code != "10000" {
		return false, errors.New(item.Message)
	}

	if item.Data.State != "1" {
		return false, errors.New("信息不一致")
	}

	return true, nil
}
