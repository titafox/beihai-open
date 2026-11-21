package routes

import (
	"coltstail-server/conf"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

var engine = newEngine()

func newEngine() *gin.Engine {
	if conf.Conf().ENV == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	e := gin.Default()
	// 只信任指定的反向代理（逗号分隔，如 "127.0.0.1,172.18.0.0/16"），
	// 否则 ClientIP() 可被 X-Forwarded-For 伪造，导致按 IP 限频失效。
	var proxies []string
	for _, p := range strings.Split(os.Getenv("TRUSTED_PROXIES"), ",") {
		if p = strings.TrimSpace(p); p != "" {
			proxies = append(proxies, p)
		}
	}
	if err := e.SetTrustedProxies(proxies); err != nil {
		log.Fatalf("invalid TRUSTED_PROXIES: %v", err)
	}
	return e
}

func New() *gin.Engine {
	return engine
}
