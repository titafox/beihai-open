package models

import (
	"errors"
	"fmt"
	"time"
)

const (
	SmsCodeTTL      = 5 * 60           // 验证码有效期（秒）
	SmsMaxFailures  = 5                // 同一手机号最多连续输错次数
	SmsFailLockTime = 15 * time.Minute // 输错过多后的锁定时间
)

type TblSms struct {
	ID       uint   `json:"id"`
	Mobile   string `json:"mobile"`
	Code     string `json:"code"`
	Event    int    `json:"event"`
	SendTime int64  `json:"send_time"`
}

func (sms *TblSms) Insert() error {
	var smsDb = Eloquent
	return smsDb.Create(sms).Error
}

// 限制2分钟发一次
func (sms *TblSms) FindSendCount(mobile string, event int) error {
	var smsDb = Eloquent
	err := smsDb.Model(TblSms{}).Where("mobile = ? AND event = ?", mobile, event).Order("id DESC").First(sms).Error
	if err == nil && time.Now().Unix()-sms.SendTime < 120 {
		return errors.New("操作频繁，请稍后再试。")
	}
	return nil
}

func smsFailKey(mobile string, event int) string {
	return fmt.Sprintf("sms:fail:%d:%s", event, mobile)
}

// CheckSmsCode 校验验证码：
//   - 只认该手机号最近一次发送的验证码，且必须在有效期内
//   - 连续输错 SmsMaxFailures 次后锁定 SmsFailLockTime
//   - 校验成功后立即作废（删除该手机号该事件的全部验证码），不可复用
func (sms *TblSms) CheckSmsCode(mobile, code string, event int) error {
	failKey := smsFailKey(mobile, event)
	if Redis != nil {
		if n, _ := Redis.Get(failKey).Int(); n >= SmsMaxFailures {
			return errors.New("验证码错误次数过多，请稍后再试")
		}
	}

	var smsDb = Eloquent
	err := smsDb.Model(TblSms{}).Where("mobile = ? AND event = ?", mobile, event).Order("id DESC").First(sms).Error
	if err != nil || code == "" || sms.Code != code || time.Now().Unix()-sms.SendTime > SmsCodeTTL {
		if Redis != nil {
			n, _ := Redis.Incr(failKey).Result()
			if n == 1 {
				Redis.Expire(failKey, SmsFailLockTime)
			}
		}
		return errors.New("验证码错误或已过期")
	}

	// 一次性使用
	smsDb.Where("mobile = ? AND event = ?", mobile, event).Delete(TblSms{})
	if Redis != nil {
		Redis.Del(failKey)
	}
	return nil
}
