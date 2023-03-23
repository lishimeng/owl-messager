package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	ses "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/ses/v20201002"
)

func Init(appId, secret string) (client *ses.Client, err error) {
	credential := common.NewCredential(appId, secret)
	client, err = ses.NewClient(credential, regions.Shanghai, profile.NewClientProfile())
	return
}
