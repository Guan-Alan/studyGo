package produce_api

import (
	"context"

	"github.com/carlmjohnson/requests"
)

// Service 公用节点
var Service = CallService{}

// CallService Produce-APi调用Service
// 调用api服务是走的http协议
type CallService struct{}

type ProduceApiConfig struct {
	Url   string `json:"url"`
	Token string `json:"token"`
}

var config ProduceApiConfig

func init() {
	config.Url = "http://test-produce.aplum-inc.com/produce-api"
	config.Token = "1111"
	return
}

type Resp struct {
	Data interface{} `json:"data"`
	Code int64       `json:"code"`
	Msg  string      `json:"msg"`
}

// PlatformWashReceive 清洗维修收货
func (s *CallService) PlatformWashReceive(ctx context.Context, washId int64) error {
	req := map[string]interface{}{"id": washId}
	resp := Resp{}
	if err := requests.URL("http://127.0.0.1:9902/produce-api/v1/platform-wash/receive").Cookie("HBL_TOKEN", config.Token).
		BodyJSON(&req).ToJSON(&resp).Fetch(ctx); err != nil {
		return err
	}
	return nil
}
