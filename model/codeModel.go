package model

/**
用户提交代码结构体  以字节数组格式加入消息队列 消费时转为json/结构体
*/
type CodeModel struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	Type string `json:"type"`
	PId  string `json:"pId"`
}
