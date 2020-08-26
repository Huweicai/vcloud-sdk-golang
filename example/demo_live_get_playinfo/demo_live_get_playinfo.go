package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
	"github.com/TTvcloud/vcloud-sdk-golang/service/live"
)

func main() {
	// call below method if you dont set ak and sk in ～/.vcloud/config
	instance := live.NewInstance()
	instance.SetCredential(base.Credentials{
		AccessKeyID:     "your ak",
		SecretAccessKey: "your sk",
	})

	// or set ak and ak as follow
	//vod.NewInstance().SetAccessKey("")
	//vod.NewInstance().SetSecretKey("")
	time.Sleep(2 * time.Second)
	did := "20200723"
	ret, err := instance.MGetStreamsPlayInfo(&live.MGetStreamsPlayInfoRequest{
		Streams:            []string{"stream-107072606641062029"},
		EnableSSL:          false,
		IsCustomizedStream: false,
		EnableStreamData:   true,
		ClientInfo:         &live.ClientInfo{DeviceId: &did},
	})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	retString, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(retString))
	return
}
