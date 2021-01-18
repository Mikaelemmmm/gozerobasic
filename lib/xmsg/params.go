package xmsg

type SendSingleMsgParams struct {
	//必填
	PhoneNumbers string            //手机号码
	Content      map[string]string //发送内容
	SignName     string            //短信签名
	TemplateCode string            //模版code
	//选填
	OutId string //业务id
}
