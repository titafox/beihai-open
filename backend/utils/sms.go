package utils

import (
	"os"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendMsCode(code, mobile string) (bool, error) {
	const path = "https://dfsns.market.alicloudapi.com/data/send_sms"
	values := url.Values{}
	values.Add("content", fmt.Sprintf("code:%s", code))
	values.Add("phone_number", mobile)
	values.Add("template_id", "TPL_09828")
	reader := strings.NewReader(values.Encode())
	req, _ := http.NewRequest(http.MethodPost, path, reader)
	req.Header.Add("Authorization", "APPCODE "+os.Getenv("ALIYUN_MARKET_APPCODE"))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	var item struct {
		RequestId string `json:"request_id"`
		Status    string `json:"status"`
		Reason    string `json:"reason"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&item); err != nil {
		return false, err
	}
	if item.Status != "OK" {
		return false, fmt.Errorf("send mscode failed: %s", item.Reason)
	}
	return true, nil
}
