package port

import (
	"encoding/json"
)

type PortInfoStat struct {
	Id    string `json:"id"`
	Proto string `json:"proto"`
	RecvQ uint64 `json:"recvQ"` // 接受队列
	SendQ uint64 `json:"sendQ"` // 发送队列
	State string `json:"state"` // 端口状态
}

func (p PortInfoStat) String() string {
	s, _ := json.Marshal(p)
	return string(s)
}
