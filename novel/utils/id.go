/**
* 作者：刘时明
* 时间：2019/10/20-14:49
* 作用：雪花算法生成分布式ID
 */
package utils

import (
	"errors"
	"sync"
	"time"
)

const (
	// 机器ID占用位数
	WORKER_ID_BITS = 5
	// 数据标识ID占用位数
	DATACENTER_ID_BITS = 5
	// 序列占用位数
	SEQUENCE_BITS = 12
	// 最大支持的机器ID
	MAX_WORKER_ID = -1 ^ (-1 << WORKER_ID_BITS)
	// 最大支持的数据标识ID
	MAX_DATACENTER_ID = -1 ^ (-1 << DATACENTER_ID_BITS)
	// 序列最大值
	MAX_SEQUENCE = -1 ^ (-1 << SEQUENCE_BITS)
	// 机器ID左移位数
	WORKER_ID_SHIFT = SEQUENCE_BITS
	// 数据标识ID左移位数
	DATACENTER_ID_SHIFT = SEQUENCE_BITS + WORKER_ID_BITS
	// 时间戳左移位数
	NOW__SHIFT = SEQUENCE_BITS + DATACENTER_ID_BITS + WORKER_ID_BITS
)

type SnowflakeIdWorker struct {
	// 机器ID(0-31)
	workerId int64
	// 数据标识ID(0-31)
	dataCenterId int64
	// 上次生成的时间戳
	lastTimestamp int64
	// 锁对象
	lock *sync.Mutex
	// 毫秒内序列
	sequence int64
}

// 获取实例
func NewSnowflake(workerId, dataCenterId int64) *SnowflakeIdWorker {
	if workerId > MAX_WORKER_ID || workerId < 0 {
		panic("workerId error")
	}
	if dataCenterId > MAX_DATACENTER_ID || dataCenterId < 0 {
		panic("dataCenterId error")
	}
	return &SnowflakeIdWorker{workerId, dataCenterId, -1, &sync.Mutex{}, 0}
}

func (s *SnowflakeIdWorker) NextID() (int64, error) {
	s.lock.Lock()
	// 13位时间戳
	now := time.Now().UnixNano() / 1000
	if now < s.lastTimestamp {
		// 时钟回调抛出异常
		return -1, errors.New("nowTimestamp less lastTimestamp")
	}
	if now == s.lastTimestamp {
		s.sequence = (s.sequence + 1) & MAX_SEQUENCE
		if s.sequence == 0 {
			// 此时说明毫秒内序列溢出，阻塞到下一毫秒
			now = tilNextMillis(s.lastTimestamp)
		}
	} else {
		// 时间戳不同，序列置0
		s.sequence = 0
	}
	s.lastTimestamp = now
	defer s.lock.Unlock()
	return (now << NOW__SHIFT) | (s.dataCenterId << DATACENTER_ID_SHIFT) | (s.workerId << WORKER_ID_SHIFT) | s.sequence, nil
}

// 获取下一毫秒
func tilNextMillis(lastTimestamp int64) int64 {
	timestamp := time.Now().UnixNano() / 1000
	for timestamp <= lastTimestamp {
		timestamp = time.Now().UnixNano() / 1000
	}
	return timestamp
}
