package tests

import (
	"coltstail-server/models"
	"testing"
	"time"
)

func TestMq(t *testing.T) {

	models.QueueInit()
	//启动列队
	r := models.RabbitMQ{Conn: models.OrderConn}
	go r.MintNft()

	var s models.RabbitMQ
	err := s.Publish("mint", "333")
	t.Log(err)

	time.Sleep(10 * time.Second)
}
