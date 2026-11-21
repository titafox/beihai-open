package conf

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

// 所有配置均从环境变量读取，示例见仓库根目录的 backend/.env.example。
// 代码中不保存任何密钥。ENV=prod 时缺少必填项会直接退出。

var ENV = os.Getenv("ENV")
var loadConfigOnce sync.Once

type Config struct {
	ENV string

	// server
	HOST string
	PORT int
	URL  string

	// jwt
	JWT_ISSUER string
	JWT_SECRET string

	// db
	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_DATABASE string
	DB_CHARSET  string
	DB_CONN_STR string

	//rabbmitMQ
	RABBMIT_MQ_MINT_URL string

	//redis
	REDIS_ADDR string

	//企账通支付
	MER_CUST_ID              string //商户号
	PFX_FILE_NAME            string // 加签证书
	PFX_FILE_PWD             string //加签证书秘钥
	UNIFIED_ORDER_URL        string //统一下单url
	CREATE_WALLET_URL        string //开通钱包
	CREATE_WALLET_STATUS_URL string //开通钱包状态查询
	SELECT_WALLET_URL        string //查询钱包
	RET_URL                  string //前端支付返回
	BG_RET_URL               string //统一下单回调
	WALLET_BG_RET_URL        string //开户下单回调
	BALANCE_BG_RET_URL       string //余额支付回调(二级市场)
	MAKE_SIGN_URL            string //加签接口
	VERIFY_SIGN_URL          string //解签接口
	CERT_FILE                string //解签证书
	DivAcctId                string
	DivCustId                string

	//avata 链服务
	AVATA_URL        string //接口地址
	AVATA_API_KEY    string //接口key
	AVATA_API_SECRET string //接口秘钥
	AVATA_CLASS_ID   string //NFT类别
}

var config Config

func Conf() Config {
	loadConfigOnce.Do(loadConfig)
	return config
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
		log.Printf("invalid int for %s: %q, using %d", key, v, def)
	}
	return def
}

func loadConfig() {
	if ENV == "" {
		ENV = "dev"
	}
	log.Printf("Using %s environment, loading config from env...", ENV)

	// 汇付接口：测试环境用 testpnr，生产用 cloudpnr
	hfHost := "https://hfpay.testpnr.com"
	avataURL := "https://stage.apis.avata.bianjie.ai"
	if ENV == "prod" {
		hfHost = "https://hfpay.cloudpnr.com"
		avataURL = "https://apis.avata.bianjie.ai"
	}

	config = Config{
		ENV:  ENV,
		HOST: env("HOST", ""),
		PORT: envInt("PORT", 8000),

		JWT_ISSUER: env("JWT_ISSUER", ENV+"_issuer"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),

		DB_HOST:     env("DB_HOST", "127.0.0.1"),
		DB_PORT:     envInt("DB_PORT", 3306),
		DB_USER:     env("DB_USER", "beihai"),
		DB_PASSWORD: os.Getenv("DB_PASSWORD"),
		DB_DATABASE: env("DB_DATABASE", "beihai"),
		DB_CHARSET:  env("DB_CHARSET", "utf8mb4"),

		REDIS_ADDR:          env("REDIS_ADDR", "127.0.0.1:6379"),
		RABBMIT_MQ_MINT_URL: os.Getenv("RABBITMQ_URL"),

		MER_CUST_ID:              os.Getenv("HF_MER_CUST_ID"),
		PFX_FILE_NAME:            os.Getenv("HF_PFX_FILE"),
		PFX_FILE_PWD:             os.Getenv("HF_PFX_PASSWORD"),
		UNIFIED_ORDER_URL:        env("HF_UNIFIED_ORDER_URL", hfHost+"/api/hfpwallet/pay033"),
		CREATE_WALLET_URL:        env("HF_CREATE_WALLET_URL", hfHost+"/api/hfpwallet/w00003"),
		CREATE_WALLET_STATUS_URL: env("HF_CREATE_WALLET_STATUS_URL", hfHost+"/api/alse/qry016"),
		SELECT_WALLET_URL:        env("HF_SELECT_WALLET_URL", hfHost+"/api/hfpwallet/w00004"),
		RET_URL:                  os.Getenv("HF_RET_URL"),
		BG_RET_URL:               os.Getenv("HF_BG_RET_URL"),
		WALLET_BG_RET_URL:        os.Getenv("HF_WALLET_BG_RET_URL"),
		BALANCE_BG_RET_URL:       os.Getenv("HF_BALANCE_BG_RET_URL"),
		MAKE_SIGN_URL:            env("HF_MAKE_SIGN_URL", "http://127.0.0.1:8080/hfpcfca/cfca/makeSign"),
		VERIFY_SIGN_URL:          env("HF_VERIFY_SIGN_URL", "http://127.0.0.1:8080/hfpcfca/cfca/verifySign"),
		CERT_FILE:                os.Getenv("HF_CERT_FILE"),
		DivAcctId:                os.Getenv("HF_DIV_ACCT_ID"),
		DivCustId:                os.Getenv("HF_DIV_CUST_ID"),

		AVATA_URL:        env("AVATA_URL", avataURL),
		AVATA_API_KEY:    os.Getenv("AVATA_API_KEY"),
		AVATA_API_SECRET: os.Getenv("AVATA_API_SECRET"),
		AVATA_CLASS_ID:   os.Getenv("AVATA_CLASS_ID"),
	}

	if config.JWT_SECRET == "" || len(config.JWT_SECRET) < 32 {
		if ENV == "prod" {
			log.Fatal("JWT_SECRET must be set (>= 32 chars) in prod")
		}
		log.Println("WARNING: JWT_SECRET is empty or short; set a random value of >= 32 chars")
	}
	if ENV == "prod" {
		for k, v := range map[string]string{
			"DB_PASSWORD":      config.DB_PASSWORD,
			"RABBITMQ_URL":     config.RABBMIT_MQ_MINT_URL,
			"AVATA_API_KEY":    config.AVATA_API_KEY,
			"AVATA_API_SECRET": config.AVATA_API_SECRET,
		} {
			if v == "" {
				log.Fatalf("%s must be set in prod", k)
			}
		}
	}

	config.URL = fmt.Sprintf("%s:%d", config.HOST, config.PORT)
	config.DB_CONN_STR = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		config.DB_USER,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_DATABASE,
		config.DB_CHARSET,
	)
}
