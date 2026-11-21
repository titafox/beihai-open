package utils

import (
	"os"
	"coltstail-server/entity/res"
	"github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"time"
)

func Token() (res.UploadToken, error) {
	var token res.UploadToken
	appId := os.Getenv("COS_APPID")
	bucket := os.Getenv("COS_STS_BUCKET")

	c := sts.NewClient(
		os.Getenv("COS_SECRET_ID"),
		os.Getenv("COS_SECRET_KEY"),
		nil,
	)

	opt := &sts.CredentialOptions{
		DurationSeconds: int64(time.Hour.Seconds()),
		Region:          os.Getenv("COS_STS_REGION"),
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					Action: []string{
						"name/cos:PostObject",
						"name/cos:PutObject",
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						"qcs::cos:" + os.Getenv("COS_STS_REGION") + ":uid/" + appId + ":" + bucket + "/*",
					},
				},
			},
		},
	}

	res, err := c.GetCredential(opt)

	if err != nil {
		return token, err
	}
	now := time.Now().Unix()
	token.SessionToken = res.Credentials.SessionToken
	token.TmpSecretID = res.Credentials.TmpSecretID
	token.TmpSecretKey = res.Credentials.TmpSecretKey
	token.StartTime = now
	token.ExpiredTime = now + 9000
	return token, err

	/*
		fmt.Printf("====================================")
		resp, err := c.RequestCredential(opt)
		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		bs, _ := ioutil.ReadAll(resp.Body)
		fmt.Printf("body:%v\n", string(bs))
	*/
}
