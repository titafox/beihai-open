package tests

import (
	"coltstail-server/middlewares"
	"coltstail-server/models"
	"coltstail-server/utils"
	"testing"
	"time"
)

func TestItem(t *testing.T) {
	info, err := middlewares.ParseJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VySUQiOjEsInVzZXJuYW1lIjoiMTg1NDk4MTIzMDQiLCJSb2xlSUQiOjAsImV4cCI6MTY2NTMwODI1MywiaXNzIjoiZGV2X2lzc3VlciIsIm5iZiI6MTY2NTIyMTg1M30.Qv7E8XTifalIXI0pGiCY1iLAJ76WOnzHIec3GwjjOr0")
	t.Log(info, err)
}

func TestGenerateRandomString(t *testing.T) {
	s := utils.GenerateRandomString(6)
	t.Log(s)
}

func TestGetNfts(t *testing.T) {
	l, s := utils.GetNftAll("0x68caf7ff5B57ad9985E03Ddb8951AE2c9Fc7E138", "0xF1Db72B50735C3E7Dc6EA69A0390f0464D63872E", "mumbai", "50")
	t.Log(l, s)
}

func TestGetGameNft(t *testing.T) {
	l, s := models.GetGameNft("0xF1Db72B50735C3E7Dc6EA69A0390f0464D63872E", "0x68caf7ff5B57ad9985E03Ddb8951AE2c9Fc7E138", 1)
	t.Log(l, s)
}

func TestUpdate(t *testing.T) {
	err := models.UpdateGameNftById(2, map[string]interface{}{"owner_address": "0xAab9645C46991Ef576910E44fdB8CA430039615e", "status": 2, "updated_at": time.Now()})
	t.Log(err)
}
