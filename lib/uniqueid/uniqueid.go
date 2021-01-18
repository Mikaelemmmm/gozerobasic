package uniqueid

import (
	"github.com/sony/sonyflake"
	"github.com/tal-tech/go-zero/core/logx"
)

var flake *sonyflake.Sonyflake

func init()  {
	flake = sonyflake.NewSonyflake(sonyflake.Settings{})
}

func GenId() int64 {

	id, err := flake.NextID()
	if err != nil {
		logx.Errorf("flake.NextID() failed with %s\n", err)
		panic(err)
	}
	return int64(id)
}
