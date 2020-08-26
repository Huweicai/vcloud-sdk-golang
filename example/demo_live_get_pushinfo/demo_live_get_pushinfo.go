package main

import (
	"encoding/json"
	"fmt"

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
	ret, err := instance.MGetStreamsPushInfo(&live.MGetStreamsPushInfoRequest{Streams: []string{"stream-107072606641062029"}})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ret.Result.PushInfos["stream-107072606641062029"].Main.RtmpUrl)

	retString, err := json.Marshal(ret)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(retString))
	return
}
