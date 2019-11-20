package imagex

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/TTvcloud/vcloud-sdk-golang/base"
)

const (
	ImageXHostCn = "imagex.bytedanceapi.com"
	ImageXHostVa = "imagex.us-east-1.bytedanceapi.com"
	ImageXHostSg = "imagex.ap-singapore-1.bytedanceapi.com"

	ImageXTimeout     = 5 * time.Second
	ImageXServiceName = "ImageX"
	ImageXApiVersion  = "2018-08-01"
)

type ImageXClient struct {
	*base.Client
}

func NewInstance() *ImageXClient {
	instance := &ImageXClient{
		Client: base.NewClient(ServiceInfoMap[base.RegionCnNorth1], ApiInfoList),
	}
	return instance
}

func NewInstanceWithRegion(region string) *ImageXClient {
	serviceInfo, ok := ServiceInfoMap[region]
	if !ok {
		panic(fmt.Errorf("can't find region %s, please check it carefully", region))
	}
	instance := &ImageXClient{
		Client: base.NewClient(serviceInfo, ApiInfoList),
	}
	return instance
}

var (
	ServiceInfoMap = map[string]*base.ServiceInfo{
		base.RegionCnNorth1: {
			Timeout: ImageXTimeout,
			Host:    ImageXHostCn,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionCnNorth1,
				Service: ImageXServiceName,
			},
		},
		base.RegionUsEast1: {
			Timeout: ImageXTimeout,
			Host:    ImageXHostVa,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionUsEast1,
				Service: ImageXServiceName,
			},
		},
		base.RegionApSingapore: {
			Timeout: ImageXTimeout,
			Host:    ImageXHostSg,
			Header: http.Header{
				"Accept": []string{"application/json"},
			},
			Credentials: base.Credentials{
				Region:  base.RegionApSingapore,
				Service: ImageXServiceName,
			},
		},
	}

	ApiInfoList = map[string]*base.ApiInfo{
		// 资源管理相关
		"ApplyUploadImageFile": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"ApplyUploadImageFile"},
				"Version": []string{ImageXApiVersion},
			},
		},
		"CommitUploadImageFile": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"CommitUploadImageFile"},
				"Version": []string{ImageXApiVersion},
			},
		},

		// 模板相关
		"GetImageTemplateConf": {
			Method: http.MethodGet,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"GetImageTemplateConf"},
				"Version": []string{ImageXApiVersion},
			},
		},
	}
)
