package tests

import (
	"coltstail-server/models"
	"testing"
	"time"
)

func TestSms(t *testing.T) {
	var sms models.TblSms
	sms.Code = "8888"
	sms.Mobile = "13800000001"
	sms.Event = 1
	sms.SendTime = time.Now().Unix()
	err := sms.Insert()
	t.Log(err)
}
