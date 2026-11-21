package controllers

import (
	"coltstail-server/ecodes"
	"coltstail-server/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HeroSliderData struct {
	Title1      string `json:"title_1"`
	Title2      string `json:"title_2"`
	Title3      string `json:"title_3"`
	Description string `json:"description"`
	Img         string `json:"img"`
	Imgbg       string `json:"imgbg"`
	Class       string `json:"class"`
}

func HeroSlider(ctx *gin.Context) {
	//data := bo.Carousel()
	data := make([]HeroSliderData, 3)
	data[0].Title1 = "Discover, find,fuck"
	data[0].Title2 = "Sell extraordinary"
	data[0].Title3 = "Monster NFTs"
	data[0].Description = "发现、收集和销售非凡的NFT"
	data[0].Img = "../assets/images/slider/slide_1.png"
	data[0].Imgbg = "https://demothemesflat.com/axies/assets/images/box-item/imgslider-3.png"
	data[0].Class = "eft"

	data[1].Title1 = "Discover, find,"
	data[1].Title2 = "Sell extraordinary"
	data[1].Title3 = "Monster NFTs"
	data[1].Description = "这是个测试"
	data[1].Img = ""
	data[1].Imgbg = ""
	data[1].Class = "center"
	//
	data[2].Title1 = "Discover, find,"
	data[2].Title2 = "Sell extraordinary"
	data[2].Title3 = "Monster NFTs"
	data[2].Description = "这是个测试"
	data[2].Img = "../assets/images/slider/slide_1.png"
	data[2].Imgbg = "https://demothemesflat.com/axies/assets/images/box-item/imgslider-3.png"
	data[2].Class = "right"

	ServerResponse(ctx, http.StatusOK, ecodes.SUCCESS, "", data)
}

func RabbmitTest(ctx *gin.Context) {
	r := models.RabbitMQ{Conn: models.OrderConn}
	err := r.Publish("mint", "2222")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
		"data":    err,
		"success": true,
	})

}
