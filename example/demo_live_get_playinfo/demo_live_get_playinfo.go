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
	ret, err := instance.MGetStreamsPlayInfo(&live.MGetStreamsPlayInfoRequest{
		Streams:            []string{"stream-106121422448623747"},
		EnableSSL:          false,
		IsCustomizedStream: false,
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
