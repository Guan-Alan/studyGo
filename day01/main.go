package main

import (
	"go.uber.org/zap"
	"guanyaowen.com/studyGo/logger"
)

type ConsumerMsg struct {
	MessageId     string `protobuf:"bytes,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`             // 消息ID
	MessageKey    string `protobuf:"bytes,2,opt,name=message_key,json=messageKey,proto3" json:"message_key,omitempty"`          // 消息key
	MessageBody   string `protobuf:"bytes,3,opt,name=message_body,json=messageBody,proto3" json:"message_body,omitempty"`       // 消息体
	ReceiptHandle string `protobuf:"bytes,4,opt,name=receipt_handle,json=receiptHandle,proto3" json:"receipt_handle,omitempty"` // 消息句柄用于ack
}

func main() {
	data := []*ConsumerMsg{
		{MessageId: "aaaaaaa", MessageBody: "adsfsdafsdfdsfa"},
		{MessageId: "bbbbbbb", MessageBody: "{\"product_id\":11111285,\"user_id\":1359,\"remark\":\"\\u5bc4\\u56de\\u6da6\\u6ed1\\u62c9\\u94fe\",\"wash_type\":\"repair\",\"biz_type\":\"aftersale_wash\",\"return_warehouse_id\":3,\"return_area_id\":53,\"double_write\":1}"},
	}
	logger.Info("msg", zap.Any("msg", data))
}
