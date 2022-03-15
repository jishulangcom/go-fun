package fun

import (
	"errors"
	"sync"
	"time"
)

/*
*************************************
--------------技术狼------------------
-------------常用方法------------------
-------------雪花算法------------------
*************************************
*/

const (
	workerBits  uint8 = 10
	numberBits  uint8 = 12
	workerMax   int64 = -1 ^ (-1 << workerBits)
	numberMax   int64 = -1 ^ (-1 << numberBits)
	timeShift   uint8 = workerBits + numberBits
	workerShift uint8 = numberBits
	startTime   int64 = 1525705533000 // 如果在程序跑了一段时间修改了epoch这个值 可能会导致生成相同的ID
)

type snowflakeWorker struct {
	mu        sync.Mutex
	timestamp int64
	workerId  int64
	number    int64
}

/*
	【名称:】创建雪花算法节点
	【参数:】
	【返回:】int64
	【备注:】
*/
func Snowflake(workerId int64) (*snowflakeWorker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New("Worker ID 超出数量")
	}
	// 生成一个新节点
	return &snowflakeWorker{
		timestamp: 0,
		workerId:  workerId,
		number:    0,
	}, nil
}

/*
	【名称:】获取雪花ID
	【参数:】
	【返回:】int64
	【备注:】
*/
func (w *snowflakeWorker) GetId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.number = 0
		w.timestamp = now
	}
	ID := int64((now-startTime)<<timeShift | (w.workerId << workerShift) | (w.number))
	return ID
}