package model

import (
	"github.com/tal-tech/go-zero/core/stringx"
	"testing"
)

/**
* @Description TODO
* @Author Mikael
* @Date 2021/1/8 23:14
* @Version 1.0
**/

func TestAa(t *testing.T)  {

	t.Log(stringx.Remove(userFieldNames, "create_time", "update_time","birthday"))
}