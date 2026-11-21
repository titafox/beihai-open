package routes

import (
	"coltstail-server/conf"
	"coltstail-server/controllers"
	_ "coltstail-server/docs"
	"coltstail-server/middlewares"
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	fmt.Println("init route")
	engine.Use(middlewares.Cors())
	// Swagger 文档只在非生产环境开放
	if conf.Conf().ENV != "prod" {
		engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//北海接口
	bhRouter := engine.Group("/api/v1")

	//版本更新
	bhRouter.GET("/app/version", controllers.GetAppVersion)
	//获取轮播图
	bhRouter.GET("/banner/list", controllers.GetBannerList)

	/*公告相关*/
	//公告类型
	bhRouter.GET("/notice/type", controllers.GetNoticeType)
	//综合公告
	bhRouter.GET("/notice/type/all", controllers.GetNoticeAll)
	//某个分类下的公告
	bhRouter.GET("/notice/type/list/:nid", controllers.GetTypeNoticeAll)
	//小喇叭
	bhRouter.GET("/notice", controllers.GetNoticeLaba)
	//详情
	bhRouter.GET("/notice/detail/:id", controllers.GetNoticeDetail)

	/*用户相关*/
	//获取短信验证码
	bhRouter.POST("/account/smsLogin/code", controllers.Send)
	//短信验证登录
	bhRouter.POST("/account/smsLogin", controllers.Login)
	//实名认证
	bhRouter.POST("/account/authentication", middlewares.AuthenticateJWT(), controllers.AccountAuthentication)
	//上传个人图像
	bhRouter.POST("/storage/avatarUpload", middlewares.AuthenticateJWT(), controllers.AvatarUpload)
	//更新用户信息
	bhRouter.POST("/account/update", middlewares.AuthenticateJWT(), controllers.UpdateAccountInfo)
	// 修改手机号
	// /account/updatePhone 原先直接映射到 Login，等于第二个登录入口，已移除
	//获取用户信息
	bhRouter.GET("/account/info", middlewares.AuthenticateJWT(), controllers.GetAccountInfo)
	//bsn二维码 todo
	bhRouter.GET("/account/qrcode", controllers.GetQrCode)
	bhRouter.GET("/account/artists", controllers.ArtList)

	/*藏品相关*/
	//藏品列表
	bhRouter.GET("/product/list", middlewares.AuthenticateJWT(), controllers.GetWorkList)
	// 作品详情
	bhRouter.GET("/product/detail/:id", middlewares.AuthenticateJWT(), controllers.GetWorkInfo)

	//我的藏品
	bhRouter.GET("/product/buyList", middlewares.AuthenticateJWT(), controllers.GetMyWorkList)
	//我的藏品详情
	bhRouter.GET("/product/buyList/:id", middlewares.AuthenticateJWT(), controllers.GetMyNftDetail)

	/*订单相关*/
	//下单
	bhRouter.POST("/order/commit", middlewares.AuthenticateJWT(), controllers.Create)
	//下单回调，一级市场
	bhRouter.POST("/pay/notify", controllers.Pay)
	//下单回调，二级市场
	bhRouter.POST("/balance-pay/notify", controllers.BalancePay)
	//订单记录
	bhRouter.GET("/order/list/:typ", middlewares.AuthenticateJWT(), controllers.OrderList)

	//BSN回调
	bhRouter.POST("/bsn/notify", controllers.BsnNotify)

	//转增
	bhRouter.POST("/works/handsel", middlewares.AuthenticateJWT(), controllers.Handsel)

	//转增记录
	bhRouter.GET("/handsel/list", middlewares.AuthenticateJWT(), controllers.HandselList)

	//v2 新增 2022年11月01日16:26:26

	//钱包开户
	bhRouter.POST("/wallet/create", middlewares.AuthenticateJWT(), controllers.WalletCreate)
	//钱包开户回调
	bhRouter.POST("/wallet/notify", controllers.WalletNotify)
	//钱包信息查询
	bhRouter.GET("/wallet/info", middlewares.AuthenticateJWT(), controllers.WalletSelect)
	//market
	bhRouter.GET("/market", controllers.Market)
	//藏品详情
	bhRouter.GET("/market/detail", middlewares.AuthenticateJWT(), controllers.MarketInfo)

	//我的藏品
	bhRouter.GET("/account/nfts", middlewares.AuthenticateJWT(), controllers.MyNFTs)
	//我的藏品合集下的作品
	bhRouter.GET("/account/nfts/series", middlewares.AuthenticateJWT(), controllers.MyNFTsSeries)
	//流通列表
	bhRouter.GET("/market/nfts/circulating", controllers.GetWorksNFT)
	//我的寄售
	bhRouter.GET("/account/nfts/sells", middlewares.AuthenticateJWT(), controllers.MyNFTSells)
	//寄售
	bhRouter.POST("/market/nft/on", middlewares.AuthenticateJWT(), controllers.MarketOnTheNFT)
	//下架
	bhRouter.POST("/market/nft/off", middlewares.AuthenticateJWT(), controllers.MarketOffTheNFT)
	//下单
	bhRouter.POST("/market/nft/buy", middlewares.AuthenticateJWT(), controllers.BuyNFT)

	//我的邀请
	bhRouter.GET("/account/invite/list", middlewares.AuthenticateJWT(), controllers.MyInviteList)
	bhRouter.GET("/account/invite/top", middlewares.AuthenticateJWT(), controllers.InviteTop)

	//k线
	bhRouter.GET("/line", controllers.GetLine)
	bhRouter.GET("/done", controllers.GetDone)

	//在线人数
	bhRouter.GET("/users", controllers.GetUsers)

	//盯链
	bhRouter.GET("/dingblock", controllers.Ding)

	/*第三期 合成*/
	//合成活动列表
	bhRouter.GET("/nft/composite/list", controllers.CompositeNFT)
	//获取合成材料
	bhRouter.GET("/nft/composite/material", middlewares.AuthenticateJWT(),controllers.GetCompositeNFTMaterial)
	//开始合成
	bhRouter.POST("/nft/composite/start", middlewares.AuthenticateJWT(),controllers.CompositeStart)
}


