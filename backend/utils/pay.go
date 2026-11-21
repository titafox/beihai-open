package utils

import (
	"coltstail-server/conf"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// 加密参数
type PfxData struct {
	PfxFileName string `json:"pfx_file_name"`
	PfxFilePwd  string `json:"pfx_file_pwd"`
}

// 加签请求参数
type MakeSignFrom struct {
	Data   string `json:"data"`
	Params string `json:"params"`
}

// 解签请求参数
type VerifySignFrom struct {
	CheckValue string `json:"check_value" comment:"加密参数"`
	CertFile   string `json:"cert_file" comment:"解密数据证书服务器路径"`
}

// 加签返回加密参数
type MakeSignData struct {
	CheckValue string `json:"check_value" comment:"加密参数"`
}

type PayReturn struct {
	Params   string `json:"params"`
	RespCode string `json:"resp_code"`
	RespDesc string `json:"resp_desc"`
}

type PayReturn2 struct {
	RespCode    string `json:"resp_code"`
	RespDesc    string `json:"resp_desc"`
	PayUrl      string `json:"pay_url"`
	RedirectUrl string `json:"redirect_url"`
}

type PayNotifyReturn struct {
	RespCode   string `json:"resp_code"`
	RespDesc   string `json:"resp_desc"`
	OrderId    string `json:"order_id"`
	UserId     string `json:"user_id"`
	UserCustId string `json:"user_cust_id"`
}

type HfPayData struct {
	MerCustId   string `json:"mer_cust_id" comment:"商户号"`
	Version     string `json:"version" comment:"版本号"`
	OrderDate   string `json:"order_date" comment:"订单日期"`
	OrderId     string `json:"order_id" comment:"订单编号，唯一"`
	TransAmt    string `json:"trans_amt" comment:"订单金额，最多两位小数"`
	GoodsDesc   string `json:"goods_desc" comment:"商品描述，用作品名称+编号"`
	UserName    string `json:"user_name" comment:"真实姓名"`
	IdCardType  string `json:"id_card_type" comment:"证件类型"`
	IdCard      string `json:"id_card" comment:"身份证号"`
	InCustId    string `json:"in_cust_id" comment:"入账客户号"`
	RetUrl      string `json:"ret_url" comment:"支付成功，前端转跳页面"`
	BgRetUrl    string `json:"bg_ret_url" comment:"支付异步回调链接"`
	DevInfoJson string `json:"dev_info_json"`
	ObjectInfo  string `json:"object_info"`
	CheckValue  string `json:"check_value" comment:"加密参数"`
}

type CreateWalletParam struct {
	OrderId  string
	UserId   string
	UserName string
	IdCard   string
}

type SelectWalletParam struct {
	OrderId  string
	UserId   string
	UserName string
	IdCard   string
}
type HfCreateWalletData struct {
	Version       string `json:"version" comment:"版本号"`
	MerCustId     string `json:"mer_cust_id" comment:"商户号"`
	OrderDate     string `json:"order_date" comment:"订单日期"`
	OrderId       string `json:"order_id" comment:"订单编号，唯一"`
	UserName      string `json:"user_name" comment:"真实姓名"`
	MarketType    string `json:"market_type"`
	AcctUsageType string `json:"acct_usage_type"`
	IdCardType    string `json:"id_card_type" comment:"证件类型"`
	IdCard        string `json:"id_card" comment:"身份证号"`
	UserId        string `json:"user_id"`
	//BgRetUrl      string `json:"bg_ret_url"`
	CheckValue string `json:"check_value" comment:"加密参数"`
}

type HfCreateWalletStatusData struct {
	Version    string `json:"version" comment:"版本号"`
	MerCustId  string `json:"mer_cust_id" comment:"商户号"`
	UserId     string `json:"user_id"`
	CheckValue string `json:"check_value" comment:"加密参数"`
}

type HfSelectWalletData struct {
	Version    string `json:"version" comment:"版本号"`
	MerCustId  string `json:"mer_cust_id" comment:"商户号"`
	UserId     string `json:"user_id"`
	OrderDate  string `json:"order_date" comment:"订单日期"`
	OrderId    string `json:"order_id" comment:"订单编号，唯一"`
	CheckValue string `json:"check_value" comment:"加密参数"`
}

type HfPayResultData struct {
	MerCustId  string `json:"mer_cust_id" comment:"商户号"`
	Version    string `json:"version" comment:"版本号"`
	OrderDate  string `json:"order_date" comment:"订单日期"`
	OrderId    string `json:"order_id" comment:"订单编号，唯一"`
	TransType  string `json:"trans_type" comment:"交易类型"`
	CheckValue string `json:"check_value" comment:"加密参数"`
}

var Pfx = PfxData{PfxFileName: conf.Conf().PFX_FILE_NAME, PfxFilePwd: conf.Conf().PFX_FILE_PWD}

type PayParam struct {
	OrderId   string
	TransAmt  string
	GoodsDesc string
	UserName  string
	IdCard    string
}

type HfBalancePayData struct {
	MerCustId       string `json:"mer_cust_id" comment:"商户号"`
	Version         string `json:"version" comment:"版本号"`
	OrderDate       string `json:"order_date" comment:"订单日期"`
	OrderId         string `json:"order_id" comment:"订单编号，唯一"`
	UserCustId      string `json:"user_cust_id" comment:"用户客户号"`
	DivDetails      string `json:"div_details" comment:"分账账户串"`
	TransAmt        string `json:"trans_amt" comment:"订单金额，最多两位小数"`
	OrderExpireTime string `json:"order_expire_time" comment:"订单过期时间"`
	GoodsDesc       string `json:"goods_desc" comment:"商品描述，用作品名称+编号"`
	//	InCustId    string `json:"in_cust_id" comment:"入账客户号"`
	RetUrl      string `json:"ret_url" comment:"支付成功，前端转跳页面"`
	BgRetUrl    string `json:"bg_ret_url" comment:"支付异步回调链接"`
	DevInfoJson string `json:"dev_info_json"`
	ObjectInfo  string `json:"object_info"`
	CheckValue  string `json:"check_value" comment:"加密参数"`
}
type BalancePayParam struct {
	OrderId    string
	TransAmt   string
	GoodsDesc  string
	UserCustId string
	DivDetails string
	NotifyUrl  string
}

// 快捷支付
func (pp PayParam) Pay() (string, error) {

	ip := ClientIP(&http.Request{})
	//统一下单参数
	var hfPayData = HfPayData{
		MerCustId:   conf.Conf().MER_CUST_ID,
		Version:     "10",
		OrderDate:   time.Now().Format("20060102"),
		OrderId:     pp.OrderId,
		TransAmt:    pp.TransAmt,
		GoodsDesc:   pp.GoodsDesc,
		UserName:    pp.UserName,
		IdCardType:  "10",
		IdCard:      pp.IdCard,
		InCustId:    conf.Conf().MER_CUST_ID,
		RetUrl:      conf.Conf().RET_URL,
		BgRetUrl:    conf.Conf().BG_RET_URL, //回调post接收到的是个加密字符串，用解签方法进行解密获取内容
		DevInfoJson: `{"devType": "1","ipAddr": "` + ip + `","IMEI": "011472001976595"}`,
		ObjectInfo:  `{"marketType":"1"}`,
	}

	//将下单参数加密
	data, _ := json.Marshal(&Pfx) //秘钥
	params, _ := json.Marshal(&hfPayData)

	payload := fmt.Sprintf("data=%s&params=%s", string(data), string(params))
	body, err := Send(conf.Conf().MAKE_SIGN_URL, payload)

	if err != nil {
		fmt.Println("支付问题1：", err)
		return "", err
	}
	var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}
	//开始调用下单接口
	payload = fmt.Sprintf("mer_cust_id=%s&version=%s&order_date=%s&order_id=%s&trans_amt=%s&goods_desc=%s&dev_info_json=%s&object_info=%s&user_name=%s&id_card_type=%s&id_card=%s&in_cust_id=%s&ret_url=%s&bg_ret_url=%s&check_value=%s",
		hfPayData.MerCustId,
		hfPayData.Version,
		hfPayData.OrderDate,
		hfPayData.OrderId,
		hfPayData.TransAmt,
		hfPayData.GoodsDesc,
		hfPayData.DevInfoJson,
		hfPayData.ObjectInfo,
		hfPayData.UserName,
		hfPayData.IdCardType,
		hfPayData.IdCard,
		hfPayData.InCustId,
		hfPayData.RetUrl,
		hfPayData.BgRetUrl,
		ret.CheckValue,
	)

	body, err = Send(conf.Conf().UNIFIED_ORDER_URL, payload)
	if err != nil {
		return "", err
	}
	//var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}

	//解签，解密下单返回的参数
	var vs VerifySignFrom
	vs.CheckValue = ret.CheckValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload = fmt.Sprintf("params=%s", string(p))

	body, err = Send(conf.Conf().VERIFY_SIGN_URL, payload)
	if err != nil {
		return "", err
	}

	//t.Log(string(body))

	//获取解析数据,拿出pay_url返回给前端支付
	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	fmt.Println("trans_amt=", pp.TransAmt)
	fmt.Println("pay=", pay)
	if err != nil {
		return "", err
	}

	if pay.RespCode != "C00000" {
		return "", errors.New("支付遇到错误")
	}

	var s2 PayReturn2
	err = json.Unmarshal([]byte(pay.Params), &s2)
	if err != nil {
		return "", err
	}

	if s2.RespCode != "C00000" {
		return "", errors.New("支付遇到错误")
	}

	return s2.PayUrl, err

}

// 余额支付
func (pp BalancePayParam) BalancePay() (string, error) {

	ip := ClientIP(&http.Request{})
	//统一下单参数
	var hfPayData = HfBalancePayData{
		MerCustId:       conf.Conf().MER_CUST_ID,
		Version:         "10",
		OrderDate:       time.Now().Format("20060102"),
		OrderId:         pp.OrderId,
		UserCustId:      pp.UserCustId,
		DivDetails:      pp.DivDetails,
		TransAmt:        pp.TransAmt,
		OrderExpireTime: time.Unix(time.Now().Unix()+180, 0).Format("20060102150405"), //锁定5分钟
		GoodsDesc:       pp.GoodsDesc,
		//InCustId:    conf.Conf().MER_CUST_ID,
		RetUrl:      conf.Conf().RET_URL,
		BgRetUrl:    pp.NotifyUrl, //回调post接收到的是个加密字符串，用解签方法进行解密获取内容
		DevInfoJson: `{"devType": "1","ipAddr": "` + ip + `","IMEI": "011472001976595"}`,
		ObjectInfo:  `{"marketType":"2"}`,
	}

	//将下单参数加密
	data, _ := json.Marshal(&Pfx) //秘钥
	params, _ := json.Marshal(&hfPayData)

	payload := fmt.Sprintf("data=%s&params=%s", string(data), string(params))
	body, err := Send(conf.Conf().MAKE_SIGN_URL, payload)

	if err != nil {
		//fmt.Println("支付问题1：", err)
		return "", err
	}
	var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}
	//开始调用下单接口
	payload = fmt.Sprintf("mer_cust_id=%s&version=%s&order_date=%s&order_id=%s&user_cust_id=%s&div_details=%s&trans_amt=%s&order_expire_time=%s&goods_desc=%s&ret_url=%s&bg_ret_url=%s&dev_info_json=%s&object_info=%s&check_value=%s",
		hfPayData.MerCustId,
		hfPayData.Version,
		hfPayData.OrderDate,
		hfPayData.OrderId,
		hfPayData.UserCustId,
		hfPayData.DivDetails,
		hfPayData.TransAmt,
		hfPayData.OrderExpireTime,
		hfPayData.GoodsDesc,
		hfPayData.DevInfoJson,
		hfPayData.ObjectInfo,
		hfPayData.RetUrl,
		hfPayData.BgRetUrl,
		ret.CheckValue,
	)

	body, err = Send(conf.Conf().UNIFIED_ORDER_URL, payload)
	if err != nil {
		return "", err
	}
	//var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}

	//解签，解密下单返回的参数
	var vs VerifySignFrom
	vs.CheckValue = ret.CheckValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload = fmt.Sprintf("params=%s", string(p))

	body, err = Send(conf.Conf().VERIFY_SIGN_URL, payload)
	if err != nil {
		return "", err
	}

	//t.Log(string(body))

	//获取解析数据,拿出pay_url返回给前端支付
	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	fmt.Println("trans_amt=", pp.TransAmt)
	fmt.Println("pay=", pay)
	if err != nil {
		return "", err
	}

	if pay.RespCode != "C00000" {
		return "", errors.New(pay.RespDesc)
	}

	var s2 PayReturn2
	err = json.Unmarshal([]byte(pay.Params), &s2)
	if err != nil {
		return "", err
	}

	if s2.RespCode != "C00000" {
		return "", errors.New(s2.RespDesc)
	}

	return s2.PayUrl, err

}

// 创建汇付钱包
func (pp CreateWalletParam) CreateWallet() (string, error) {
	//开户参数
	var hfPayData = HfCreateWalletData{
		Version:       "10",
		MerCustId:     conf.Conf().MER_CUST_ID,
		OrderDate:     time.Now().Format("20060102"),
		OrderId:       pp.OrderId,
		UserName:      pp.UserName,
		MarketType:    "2",
		AcctUsageType: "wallet",
		IdCardType:    "10",
		IdCard:        pp.IdCard,
		UserId:        pp.UserId,
		//BgRetUrl:      conf.Conf().SELECT_WALLET_URL,
	}

	//将下单参数加密
	data, _ := json.Marshal(&Pfx) //秘钥
	params, _ := json.Marshal(&hfPayData)

	payload := fmt.Sprintf("data=%s&params=%s", string(data), string(params))
	body, err := Send(conf.Conf().MAKE_SIGN_URL, payload)

	if err != nil {
		fmt.Println("支付问题1：", err)
		return "", err
	}
	var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}
	//开始调用下单接口
	payload = fmt.Sprintf("version=%s&mer_cust_id=%s&order_date=%s&order_id=%s&user_name=%s&market_type=%s&acct_usage_type=%s&id_card_type=%s&id_card=%s&user_id=%s&check_value=%s",
		hfPayData.Version,
		hfPayData.MerCustId,
		hfPayData.OrderDate,
		hfPayData.OrderId,
		hfPayData.UserName,
		hfPayData.MarketType,
		hfPayData.AcctUsageType,
		hfPayData.IdCardType,
		hfPayData.IdCard,
		hfPayData.UserId,
		//hfPayData.BgRetUrl,
		ret.CheckValue,
	)

	body, err = Send(conf.Conf().CREATE_WALLET_URL, payload)
	if err != nil {
		return "", err
	}
	//var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}

	//解签，解密下单返回的参数
	var vs VerifySignFrom
	vs.CheckValue = ret.CheckValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload = fmt.Sprintf("params=%s", string(p))

	body, err = Send(conf.Conf().VERIFY_SIGN_URL, payload)
	if err != nil {
		return "", err
	}

	//t.Log(string(body))

	//fmt.Println("钱包链接：",string(body))

	//获取解析数据,拿出pay_url返回给前端支付
	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	if err != nil {
		return "", err
	}

	//fmt.Println(pay)
	if pay.RespCode != "C00000" {
		return "", errors.New(pay.RespDesc)
	}

	var s2 PayReturn2
	err = json.Unmarshal([]byte(pay.Params), &s2)
	if err != nil {
		return "", err
	}

	if s2.RespCode != "C00000" {
		return "", errors.New(s2.RespDesc)
	}

	return s2.RedirectUrl, err

}

// 汇付钱包创建状态查询
func (pp SelectWalletParam) SelectWalletStatus() (string, string, error) {
	fmt.Println("pp=", pp)
	//开会参数
	var hfPayData = HfCreateWalletStatusData{
		Version:   "10",
		MerCustId: conf.Conf().MER_CUST_ID,
		UserId:    pp.UserId,
	}

	//将下单参数加密
	data, _ := json.Marshal(&Pfx) //秘钥
	params, _ := json.Marshal(&hfPayData)

	payload := fmt.Sprintf("data=%s&params=%s", string(data), string(params))
	body, err := Send(conf.Conf().MAKE_SIGN_URL, payload)

	if err != nil {
		return "", "", err
	}
	var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", "", err
	}
	//开始查询状态
	payload = fmt.Sprintf("version=%s&mer_cust_id=%s&check_value=%s",
		hfPayData.Version,
		hfPayData.MerCustId,
		ret.CheckValue,
	)

	body, err = Send(conf.Conf().CREATE_WALLET_STATUS_URL, payload)
	if err != nil {
		return "", "", err
	}
	//var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", "", err
	}

	//解签，解密下单返回的参数
	var vs VerifySignFrom
	vs.CheckValue = ret.CheckValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload = fmt.Sprintf("params=%s", string(p))

	body, err = Send(conf.Conf().VERIFY_SIGN_URL, payload)
	if err != nil {
		return "", "", err
	}

	//t.Log(string(body))

	fmt.Println("钱包链接：", string(body))

	//获取解析数据,拿出pay_url返回给前端支付
	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	if err != nil {
		return "", "", err
	}

	if pay.RespCode != "C00000" {
		return "", "", errors.New(pay.RespDesc)
	}

	var s struct {
		RespCode     string `json:"resp_code"`
		RespDesc     string `json:"resp_desc"`
		UserCustId   string `json:"user_cust_id"`
		AcctInfoList []struct {
			AcctId string `json:"acct_id"`
		} `json:"acct_info_list,omitempty"`
	}

	err = json.Unmarshal([]byte(pay.Params), &s)
	if err != nil {
		return "", "", errors.New("未查询到开户信息")
	}

	if s.RespCode != "C00000" {
		return "", "", errors.New(s.RespDesc)
	}

	return s.UserCustId, s.AcctInfoList[0].AcctId, err

}

// 查询汇付钱包
func (pp SelectWalletParam) SelectWallet() (string, error) {
	//查询参数
	var hfPayData = HfSelectWalletData{
		Version:   "10",
		MerCustId: conf.Conf().MER_CUST_ID,
		UserId:    pp.UserId,
		OrderDate: time.Now().Format("20060102"),
		OrderId:   pp.OrderId,
	}

	//将下单参数加密
	data, _ := json.Marshal(&Pfx) //秘钥
	params, _ := json.Marshal(&hfPayData)

	payload := fmt.Sprintf("data=%s&params=%s", string(data), string(params))
	body, err := Send(conf.Conf().MAKE_SIGN_URL, payload)

	if err != nil {
		fmt.Println("支付问题1：", err)
		return "", err
	}
	var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}
	//开始调用下单接口
	payload = fmt.Sprintf("version=%s&mer_cust_id=%s&user_id=%s&check_value=%s",
		hfPayData.Version,
		hfPayData.MerCustId,
		hfPayData.UserId,
		ret.CheckValue,
	)

	body, err = Send(conf.Conf().SELECT_WALLET_URL, payload)
	if err != nil {
		return "", err
	}
	//var ret MakeSignData
	err = json.Unmarshal(body, &ret)
	if err != nil {
		return "", err
	}

	//解签，解密下单返回的参数
	var vs VerifySignFrom
	vs.CheckValue = ret.CheckValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload = fmt.Sprintf("params=%s", string(p))

	body, err = Send(conf.Conf().VERIFY_SIGN_URL, payload)
	if err != nil {
		return "", err
	}

	//t.Log(string(body))

	//fmt.Println("钱包参数：", string(body))

	//获取解析数据,拿出pay_url返回给前端支付
	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	if err != nil {
		return "", err
	}

	fmt.Println("====", pay)
	if pay.RespCode != "C00000" {
		return "", errors.New(pay.RespDesc)
	}

	var s2 PayReturn2
	err = json.Unmarshal([]byte(pay.Params), &s2)
	if err != nil {
		return "", err
	}

	if s2.RespCode != "C00000" {
		return "", errors.New(s2.RespDesc)
	}

	return s2.RedirectUrl, err

}

// 验证签名
func VerifySign(checkValue string) (string, error) {
	var vs VerifySignFrom
	vs.CheckValue = checkValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload := fmt.Sprintf("params=%s", string(p))

	body, err := Send(conf.Conf().VERIFY_SIGN_URL, payload)
	//回调解签数据
	fmt.Println("回调解签数据:", string(body), err)
	if err != nil {
		return "", err
	}

	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	if err != nil {
		return "", err
	}

	if pay.RespCode != "C00000" {
		return "", errors.New("回调错误：" + pay.RespDesc)
	}

	var s2 PayNotifyReturn
	err = json.Unmarshal([]byte(pay.Params), &s2)
	if err != nil {
		return "", err
	}

	if s2.RespCode != "C00000" {
		return "", errors.New("订单回调错误：" + s2.RespDesc)
	}

	//fmt.Println(s2.OrderId)
	return s2.OrderId, nil
}

// 钱包开户解签
func CreateWalletVerifySign(checkValue string) (string, error) {
	var vs VerifySignFrom
	vs.CheckValue = checkValue
	vs.CertFile = conf.Conf().CERT_FILE
	p, _ := json.Marshal(&vs)
	payload := fmt.Sprintf("params=%s", string(p))

	body, err := Send(conf.Conf().VERIFY_SIGN_URL, payload)
	//回调解签数据
	fmt.Println("回调解签数据:", string(body), err)
	if err != nil {
		return "", err
	}

	var pay PayReturn
	err = json.Unmarshal(body, &pay)
	if err != nil {
		return "", err
	}

	fmt.Println("开户解签数据：", pay)

	if pay.RespCode != "C00000" {
		return "", errors.New("回调错误：" + pay.RespDesc)
	}

	var s2 PayNotifyReturn
	err = json.Unmarshal([]byte(pay.Params), &s2)
	if err != nil {
		return "", err
	}

	fmt.Println("pay222222", s2)

	if s2.RespCode != "C00000" {
		return "", errors.New("订单回调错误：" + s2.RespDesc)
	}

	//fmt.Println(s2.OrderId)
	return s2.OrderId, nil
}

// 支付接口请求封装
func Send(url, payload string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	return body, err
}
