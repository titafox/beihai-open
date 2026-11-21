package utils

import (
	"bytes"
	"coltstail-server/conf"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"

	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type tag struct {
	Key1 string  `json:"key1,omitempty"` //链账户地址
	Key2 string  `json:"key2,omitempty"` //链账户名称
	Key3 float64 `json:"key3,omitempty"` //Gas 余额
}

type accountCreate struct {
	Account     string `json:"account"`      //链账户地址
	Name        string `json:"name"`         //链账户名称
	OperationId string `json:"operation_id"` //操作 ID
}

type AccountCreateReq struct {
	/*
	 *链账户名称，支持 1-20 位汉字、大小写字母及数字组成的字符串
	 */
	Name string `json:"name"`
	/*
	 *操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	 */
	OperationId string `json:"operation_id"`
}

type AccountCreateDTO struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"statusCode"`
	Data       accountCreate `json:"data"`
}

type nftCreatePathParam struct {
	ClassId string `json:"class_id"` //NFT 类别 ID
}

type nftCreateRequestBody struct {
	Name        string `json:"name"`                //NFT 名称
	Uri         string `json:"uri,omitempty"`       //链外数据链接
	UriHash     string `json:"uri_hash,omitempty"`  //链外数据 Hash
	Data        string `json:"data,omitempty"`      //自定义链上元数据
	Recipient   string `json:"recipient,omitempty"` //链外数据 Hash
	Tag         tag    `json:"tag,omitempty"`       //交易标签, 自定义 key：支持大小写英文字母和汉字和数字，长度6-12位，自定义 value：长度限制在64位字符，支持大小写字母和数字
	OperationId string `json:"operation_id"`        //操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
}

// NFTCreateReq 发行NFT
type NFTCreateReq struct {
	PathParam   nftCreatePathParam   `json:"pathParam"`   //PATH PARAMETERS
	RequestBody nftCreateRequestBody `json:"requestBody"` //REQUEST BOD
}

type nftCreate struct {
	TaskId      string `json:"task_id"`      //Task ID
	OperationId string `json:"operation_id"` //操作 ID
}

type NFTCreateDTO struct {
	Status     string    `json:"status"`
	StatusCode int       `json:"statusCode"`
	Data       nftCreate `json:"data"`
}

type nftTransferPathParam struct {
	ClassId string `json:"class_id"` //NFT 类别 ID
	Owner   string `json:"owner"`    //NFT 持有者地址
	NftId   string `json:"nft_id"`   //NFT 类别 ID
}

type nftTransferRequestBody struct {
	Recipient   string `json:"recipient"`     //NFT 接收者地址
	OperationId string `json:"operation_id"`  //操作 ID，保证幂等性，避免重复请求，保证对于同一操作发起的一次请求或者多次请求的结果是一致的；由接入方生成的、针对每个 Project ID 唯一的、不超过 64 个大小写字母、数字、-、下划线的字符串
	Tag         tag    `json:"tag,omitempty"` //交易标签, 自定义 key：支持大小写英文字母和汉字和数字，长度6-12位，自定义 value：长度限制在64位字符，支持大小写字母和数字
}

type NFTTransferReq struct {
	PathParam   nftTransferPathParam   `json:"pathParam"`   //PATH PARAMETERS
	RequestBody nftTransferRequestBody `json:"requestBody"` //REQUEST BOD
}

type nftTransfer struct {
	TaskId      string `json:"task_id"`      //Task ID
	OperationId string `json:"operation_id"` //操作 ID
}

type NFTTransferDTO struct {
	Status     string      `json:"status"`
	StatusCode int         `json:"statusCode"`
	Data       nftTransfer `json:"data"`
}

// ErrorData 错误信息
type errorData struct {
	Code      string `json:"code"`
	CodeSpace string `json:"code_space"`
	Message   string `json:"message"`
}

// ErrorDTO 错误返回值
type ErrorRes struct {
	Status     string `json:"status"`
	StatusCode int    `json:"statusCode"`
}

// ErrorDTO 错误返回值
type ErrorDTO struct {
	Err        error     `json:"err"`
	Status     string    `json:"status"`
	StatusCode int       `json:"statusCode"`
	Error      errorData `json:"error"`
}

type Avata struct {
	ClassId string `json:"class_id"`
}

// SignRequest 对请求进行签名
func SignRequest(r *http.Request, apiKey, apiSecret string) *http.Request {
	timestamp := strconv.FormatInt(time.Now().Unix()*1000, 10)
	// 获取 path params
	params := map[string]interface{}{}
	params["path_url"] = r.URL.Path
	// 获取 query params
	for k, v := range r.URL.Query() {
		k = "query_" + k
		params[k] = v[0]
	}
	// 获取 body params
	// 把request的内容读取出来
	var bodyBytes []byte
	if r.Body != nil {
		bodyBytes, _ = io.ReadAll(r.Body)
	}
	// 把刚刚读出来的再写进去
	if bodyBytes != nil {
		r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	paramsBody := map[string]interface{}{}
	_ = json.Unmarshal(bodyBytes, &paramsBody)
	hexHash := hash(timestamp + apiSecret)
	for k, v := range paramsBody {
		k = "body_" + k
		params[k] = v
	}
	sortParams := params
	if sortParams != nil {
		sortParamsBytes, _ := json.Marshal(sortParams)
		hexHash = hash(string(sortParamsBytes) + timestamp + apiSecret)
	}
	r.Header.Set("X-Api-Key", apiKey)
	r.Header.Set("X-Signature", hexHash)
	r.Header.Set("X-Timestamp", timestamp)
	return r
}
func hash(oriText string) string {
	oriTextHashBytes := sha256.Sum256([]byte(oriText))
	return hex.EncodeToString(oriTextHashBytes[:])
}

func PostHeader(url string, param []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(param))
	if err != nil {
		return nil, err
	}
	//验签
	reqS := SignRequest(req, conf.Conf().AVATA_API_KEY, conf.Conf().AVATA_API_SECRET)
	reqS.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(reqS)
	return resp, err
}

// AccountCreate 创建链账户
func AccountCreate(param *AccountCreateReq) (res AccountCreateDTO, errDTO ErrorDTO) {
	p, err := json.Marshal(&param)
	if err != nil {
		errDTO.Err = err
		return
	}
	url := conf.Conf().AVATA_URL + "/v1beta1/account"
	resp, err := PostHeader(url, p)
	if err != nil {
		errDTO.Err = err
		return
	}
	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		fmt.Println(err.Error())
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	return res, errDTO
}

// NFTCreate 发行 NFT
func NFTCreate(param *NFTCreateReq) (res NFTCreateDTO, errDTO ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}

	url := conf.Conf().AVATA_URL + "/v1beta1/nft/nfts/{class_id}"
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	fmt.Println(url)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := PostHeader(url, bytesData)
	if err != nil {
		errDTO.Err = err

		return
	}
	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("err-01:", err)
			return
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		fmt.Println("err-02:", err.Error())
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		fmt.Println("err-03:", err.Error())
		return
	}
	return res, errDTO
}

// NFTTransfer 转让 NFT
func NFTTransfer(param *NFTTransferReq) (res NFTTransferDTO, errDTO ErrorDTO) {
	if param.PathParam.ClassId == "" {
		errDTO.Err = errors.New("classId mast be provide")
		return
	}
	if param.PathParam.Owner == "" {
		errDTO.Err = errors.New("owner mast be provide")
		return
	}
	if param.PathParam.NftId == "" {
		errDTO.Err = errors.New("nft_id mast be provide")
		return
	}
	url := conf.Conf().AVATA_URL + "/v1beta1/nft/nft-transfers/{class_id}/{owner}/{nft_id}"
	url = strings.Replace(url, "{class_id}", param.PathParam.ClassId, -1)
	url = strings.Replace(url, "{owner}", param.PathParam.Owner, -1)
	url = strings.Replace(url, "{nft_id}", param.PathParam.NftId, -1)
	bytesData, err := json.Marshal(param.RequestBody)
	if err != nil {
		errDTO.Err = err
		return
	}

	resp, err := PostHeader(url, bytesData)
	if err != nil {
		errDTO.Err = err
		return
	}
	body, err := io.ReadAll(resp.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("err-08:", err.Error())
		}
	}(resp.Body)
	if err != nil {
		errDTO.Err = err
		return
	}
	if resp.StatusCode == 400 {
		err = json.Unmarshal(body, &errDTO)
	} else if resp.StatusCode == 200 {
		err = json.Unmarshal(body, &res)
	}
	// 设置状态和状态码
	errDTO.Status = resp.Status
	errDTO.StatusCode = resp.StatusCode
	res.Status = resp.Status
	res.StatusCode = resp.StatusCode
	if err != nil {
		return
	}
	return res, errDTO
}

type AccontCreateRes struct {
	StatusCode int `json:"statusCode"`
	Data       struct {
		Account     string `json:"account"`
		OperationId string `json:"operation_id"`
	} `json:"data"`
}

// 铸件NFT
func (avata *Avata) CreateAccount(name, op string) (string, error) {
	param := new(AccountCreateReq)
	param.Name = name
	param.OperationId = op
	result, err := AccountCreate(param)
	if err.Err != nil {
		return "", err.Err
	}
	byteResultStr, _ := json.Marshal(result)
	fmt.Println(string(byteResultStr))
	var caRes AccontCreateRes
	err4 := json.Unmarshal(byteResultStr, &caRes)
	if err4 != nil {
		return "", err4
	}

	if caRes.StatusCode != 200 {
		return "", errors.New("创建地址失败")
	}

	//byteErrStr, _ := json.Marshal(err)
	//logrus.Error(string(byteErrStr))

	return caRes.Data.Account, nil
}

// 创建NFT
func (avata *Avata) CreateNFT(op, recipient, name string) (bool, error) {
	param := new(NFTCreateReq)
	param.PathParam.ClassId = conf.Conf().AVATA_CLASS_ID
	param.RequestBody.Name = name
	param.RequestBody.Recipient = recipient
	//param.RequestBody.Uri = uri
	param.RequestBody.OperationId = op

	result, err := NFTCreate(param)
	fmt.Println("res=", result)
	if err.Err != nil {
		return false, err.Err
	}
	_, err2 := json.Marshal(result)
	if err2 != nil {
		return false, err2
	}
	fmt.Println("err=", err)
	byteErrStr, err3 := json.Marshal(err)
	if err3 != nil {
		return false, err3
	}
	var res ErrorDTO
	err5 := json.Unmarshal(byteErrStr, &res)
	if err5 != nil {
		return false, err5
	}

	//fmt.Println("res:", res)
	if res.StatusCode != 200 {
		return false, errors.New("创建NNT失败")
	}
	return true, nil
}

// 转让 NFT
func (avata *Avata) NftTransfer(owner, nft_id, recipient, op string) (bool, error) {
	param := new(NFTTransferReq)
	param.PathParam.ClassId = conf.Conf().AVATA_CLASS_ID
	param.PathParam.Owner = owner
	param.PathParam.NftId = nft_id
	param.RequestBody.Recipient = recipient
	param.RequestBody.OperationId = op
	_, err := NFTTransfer(param)
	if err.Err != nil {
		fmt.Println("1")
		return false, err.Err
	}
	//byteResultStr, _ := json.Marshal(result)
	//logrus.Println(string(byteResultStr))
	byteErrStr, _ := json.Marshal(err)
	var res ErrorDTO
	err5 := json.Unmarshal(byteErrStr, &res)
	if err5 != nil {
		return false, err5
	}

	if res.StatusCode != 200 {
		return false, errors.New("转移NNT失败:" + res.Error.Message)
	}
	return true, nil
}

// TxResult AVATA 上链交易结果（GET /v1beta1/tx/{operation_id}）
type TxResult struct {
	Type        string `json:"type"`
	TxHash      string `json:"tx_hash"`
	Status      int    `json:"status"` // 0 处理中 1 成功 2 失败 3 未处理
	Message     string `json:"message"`
	BlockHeight int    `json:"block_height"`
	Nft         string `json:"nft"`
}

// QueryTx 按 operation_id 向 AVATA 回查交易结果。
// 回调接口无法证明请求来自 AVATA，因此回调内容只作为"通知"，以回查结果为准。
func QueryTx(operationId string) (*TxResult, error) {
	if operationId == "" {
		return nil, errors.New("empty operation_id")
	}
	u := conf.Conf().AVATA_URL + "/v1beta1/tx/" + operationId
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}
	req = SignRequest(req, conf.Conf().AVATA_API_KEY, conf.Conf().AVATA_API_SECRET)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("avata query tx: %s %s", resp.Status, string(body))
	}
	var out struct {
		Data TxResult `json:"data"`
	}
	if err := json.Unmarshal(body, &out); err != nil {
		return nil, err
	}
	return &out.Data, nil
}
