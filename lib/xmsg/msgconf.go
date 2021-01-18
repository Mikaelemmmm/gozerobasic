package xmsg

type MsgConf struct {
	Aliyun struct {
		AccessId     string `json:",optional"`
		AccessSecret string `json:",optional"`
		RegionId     string `json:",default=cn-hangzhou"`
	}
}
