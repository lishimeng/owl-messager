package ding

import (
	"fmt"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	robot "github.com/alibabacloud-go/dingtalk/robot_1_0"
)

// 钉钉

type OtoSDK interface {
}

type otoSdkImpl struct {
	AppKey    string
	AppSecret string

	accessToken string

	proxy *robot.Client
}

func New(appKey, appSecret string) (sdk OtoSDK, err error) {

	var oto otoSdkImpl
	var client *robot.Client
	client, err = robot.NewClient(&openapi.Config{
		AccessKeyId:     &appKey,
		AccessKeySecret: &appSecret,
	})
	if err != nil {
		return
	}
	oto.proxy = client

	return
}

func (sdk *otoSdkImpl) SendRobotMessage() (err error) {
	req := robot.SendRobotDingMessageRequest{ // TODO
		ContentParams:      nil,
		DingTemplateId:     nil,
		OpenConversationId: nil,
		ReceiverUserIdList: nil,
		RobotCode:          nil,
	}

	resp, err := sdk.proxy.SendRobotDingMessage(&req)

	if err != nil {
		return
	}
	fmt.Println(resp.String())
	return
}
