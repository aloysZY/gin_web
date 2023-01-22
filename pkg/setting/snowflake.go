// Package setting
package setting

import (
	"errors"
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

// NewSonyFlake 需传入当前的机器ID,和开始时间
// startTime 是基于这个时间节点开始的增量69 年
func NewSonyFlake(machineId uint16, startTime string) (err error) {
	sonyMachineID = machineId
	t, _ := time.Parse("2006-01-02", startTime)
	settings := sonyflake.Settings{
		StartTime: t,
		MachineID: getMachineID,
	}
	sonyFlake = sonyflake.NewSonyflake(settings)
	return
}

// GetID 返回生成的id值
func GetID() (id uint64, err error) {
	// 如果没有初始化就报这个错误
	if sonyFlake == nil {
		// err = fmt.Errorf("newSonyFlake not initialized")
		return 0, errors.New("newSonyFlake not initialized")
	}
	id, err = sonyFlake.NextID()
	return
}
