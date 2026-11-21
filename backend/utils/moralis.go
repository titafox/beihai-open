package utils

import (
	"os"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Nfts struct {
	Status string `json:"status"`
	Result []struct {
		TokenId string `json:"token_id"`
	} `json:"result"`
}

// 获取某个地址合约下所有nft
func GetNftAll(address, token_address, chain, limit string) (*Nfts, error) {

	url := fmt.Sprintf("https://deep-index.moralis.io/api/v2/%s/nft/%s?chain=%s&format=decimal&limit=%s", address, token_address, chain, limit)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("x-api-key", os.Getenv("MORALIS_API_KEY"))

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var list Nfts
	if err := json.Unmarshal(body, &list); err != nil {
		return nil, err
	}

	fmt.Println(list)
	return &list, nil
}
