package main

import (
	"fmt"
	"time"

	"github.com/chenjiacheng/go-shumei"
)

func main() {
	client := shumei.NewClient(
		"Your App ID",
		"Your Access Key",
		shumei.WithRegion(shumei.RegionGuangzhou),
		shumei.WithTimeout(time.Second*5),
	)

	// 文本检测
	res, err := client.TextDetect("Your EventId", shumei.TextDetectReq{
		Type: "TEXTRISK",
		Data: map[string]interface{}{
			"text":     "加个好友吧 qq12345",
			"tokenId":  "4567898765jhgfdsa",
			"ip":       "118.89.214.89",
			"deviceId": "*************",
			"nickname": "***********",
			"extra": map[string]interface{}{
				"topic":          "12345",
				"atId":           "username1",
				"room":           "ceshi123",
				"receiveTokenId": "username2",
				"role":           "USER",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Code:", res.Code)
	fmt.Println("Message:", res.Message)
	fmt.Println("RequestID:", res.RequestID)
}
