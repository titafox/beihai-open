package main

import (
	"coltstail-server/conf"
	"coltstail-server/models"
	"coltstail-server/routes"
	"coltstail-server/services"
	"github.com/beego/beego/v2/task"
	"log"
)

// @title Coltstail APIs
// @version 1.0
// @description Coltstail APIs
// @termsOfService http://swagger.io/terms/

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8000
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api/v1

func main() {
	//定时任务
	services.InitTask()
	task.StartTask()
	defer task.StopTask()

	//go动以太坊节点监听事件
	//go bo.ListenEthNode()

	//监听马蹄事件
	//go bo.ListenMetNode()

	//启动redis
	models.InitRedis()

	//启动mq
	models.QueueInit()
	//启动列队
	r := models.RabbitMQ{Conn: models.OrderConn}
	go r.MintNft()

	//启动主程序
	config := conf.Conf()
	engine := routes.New()
	//
	////载入模版
	//engine.Delims("{<", ">}")
	//engine.Static("/static/", "./static")
	//engine.LoadHTMLGlob("static/web/templates/*")
	//engine.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", gin.H{"Title": "我是测试", "ce": "123456"})
	//})

	log.Fatal(engine.Run(config.URL))
}
