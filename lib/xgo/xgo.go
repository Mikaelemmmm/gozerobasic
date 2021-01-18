package xgo

import "github.com/tal-tech/go-zero/core/logx"

/**
* @Description TODO
* @Author Mikael
* @Date 2021/1/18 10:38
* @Version 1.0
**/
func Go(g func())  {
	defer func() {
		if r := recover(); r != nil {
			logx.Errorf("GO recover panic e r:%v",r)
		}

		go g()
	}()
}