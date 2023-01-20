// Package app 生成 ID 的算法，不使用自增 id，以免一下子就知道用户数量
package app

import (
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var (
	sonyFlake     *sonyflake.Sonyflake
	sonyMachineID uint16
)

func getMachineID() (uint16, error) {
	return sonyMachineID, nil
}

// NewSonyFlake 需传入当前的机器ID
func NewSonyFlake(machineId uint16) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", "2022-12-31")
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回生成的id值
func GetID() (id uint64, err error) {
	if sonyFlake == nil {
		err = fmt.Errorf("snoy flake not inited")
		return
	}
	id, err = sonyFlake.NextID()
	return
}
