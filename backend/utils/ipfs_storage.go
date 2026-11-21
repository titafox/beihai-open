package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type nftStorageRes struct {
	Ok    bool `json:"ok"`
	Value struct {
		Cid string `json:"cid"`
	} `json:"value"`
}

func UploadIpfsFile(file []byte) (string, error) {
	url := "https://api.nft.storage/upload"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewReader(file))
	if err != nil {
		return "", err
	}

	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJkaWQ6ZXRocjoweDgyMDU3MzMyN0FCZjkzOTExOGMxNTk0MDVFYUEzMDYxQTQ1NmZFMDQiLCJpc3MiOiJuZnQtc3RvcmFnZSIsImlhdCI6MTYzNDAwNzA0OTMzOSwibmFtZSI6InByb2QxIn0._-SSAlp0meE2bcqfgFrGtCvNCkn-hUdKjMgKvpqdVFg")
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	var nftRes nftStorageRes
	if err := json.Unmarshal(body, &nftRes); err != nil {
		return "", err
	}

	if !nftRes.Ok {
		return "", errors.New("upload fail")
	}

	return "ipfs://" + nftRes.Value.Cid, nil
}
