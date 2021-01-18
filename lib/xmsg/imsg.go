package xmsg

type IMsg interface {
	SendSingleMsg(param *SendSingleMsgParams) error
}

//如果切换其他服务商只需要改这里
func NewMsgInstance(c MsgConf) IMsg {
	return NewAliyunMsg(c.Aliyun.RegionId,c.Aliyun.AccessId,c.Aliyun.AccessSecret)
}
