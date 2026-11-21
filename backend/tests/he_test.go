package tests

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func TestHC(t *testing.T) {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go send(&wg)
	}

	wg.Wait()
}

func send(wg *sync.WaitGroup) {
	url := "http://127.0.0.1:8000/api/v1/nft/composite/start"
	method := "POST"

	payload := strings.NewReader(`{
  "composite_id": 1,
  "data": [
    {
      "ids": [
        12,13
      ],
      "wid": 2
    },
{
      "ids": [
        14,15
      ],
      "wid": 4
    }
  ]
}`)

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEwMDAwLCJ1c2VybmFtZSI6IjEzMzM1ODMxOTc5IiwiUm9sZUlEIjowLCJleHAiOjE2NzAxNDA0NDQsImlzcyI6ImRldl9pc3N1ZXIiLCJuYmYiOjE2Njk1MzU2NDR9.jtfSYJLV6an7GIkcbqE9dBr_ybYpKMlgA8snOf3xVGI")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
	wg.Done()
}
