package enum

/**
* @Description 用户授权枚举
* @Author Mikael
* @Date 2020/12/6 12:55
* @Version 1.0
**/

type AuthTypeEnum int64

/**
	授权类型
 */
const (
	AuthTypeMobile  AuthTypeEnum = 0
	AuthTypeSmallWx AuthTypeEnum = 1
)
func (em AuthTypeEnum)String() string {
	switch em {
	case AuthTypeMobile:
		return "MOBILE"
	case AuthTypeSmallWx:
		return "SMALL_EX"
	default:
		return "unknow auth type"
	}
}