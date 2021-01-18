package enum

import "gozerobasic/lib/xerr"

/**
* @Description 短信配置枚举
* @Author Mikael
* @Date 2020/12/6 17:40
* @Version 1.0
**/

var errMessageTplCodeNoFound = xerr.NewErrMsg("短信模版code不存在")


/**
	短信签名枚举
 */
type MessageSignEnum int64
const (
	MessageSignFishTwo  MessageSignEnum = 0
)
func (em MessageSignEnum)String() string {
	switch em {
	case MessageSignFishTwo:
		return "鱼二"
	default:
		panic("unkonw message sign")
	}
}


/**
	短信类型
 */
type MessageKindEnum int64
const (
	MessageKindRegister  MessageKindEnum = 0 //注册
	MessageKindLogin     MessageKindEnum = 1 //登陆
	MessageKindForgetPwd MessageKindEnum = 2 //忘记密码
)
func (em MessageKindEnum)String() string {
	switch em {
	case MessageKindRegister:
		return "注册"
	case MessageKindLogin:
		return "登陆"
	case MessageKindForgetPwd:
		return "忘记密码"
	default:
		panic("unkonw message kind")
	}
}
//获取短信模版code
func (em MessageKindEnum)GetMessageTplConfig() (tplCode string, tplSign string, err error) {

	switch em {
	case MessageKindRegister:
		return "SMS_162040024", MessageSignFishTwo.String(), nil
	case MessageKindLogin:
		return "SMS_162040024", MessageSignFishTwo.String(), nil
	case MessageKindForgetPwd:
		return "SMS_162040024", MessageSignFishTwo.String(), nil
	}

	return "", "", errMessageTplCodeNoFound
}
