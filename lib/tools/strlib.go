package tools

import "strings"

//隐藏手机号码中间4位
func HideMobileMid(phone string) string {

	old := ""
	for k, v := range phone{
		if k >=3 && k <=6{
			old = old + string(v)
		}
	}
	return strings.Replace(phone,old,"****",-1)
}
