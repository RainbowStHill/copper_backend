package identity_server

import (
	"sync"
	"time"
)

type SnowflakeOpt struct {
	// size of memory space storing timestamp in bit
	TimestampBits int64
	// size of memory space storing data center ID in bit
	DataCenterIDBits int64
	// size of service instance ID in bit
	InstanceIDBits int64
	// size of the naturally increasing ID in bit
	IncrIDBits int64

	// Time when this system starts.
	StartsAt   int64
	DataCenter int64
	Instance   int64
}

type snowflake struct {
	SnowflakeOpt

	timestampMaxValue    int64
	dataCenterIDMaxValue int64
	instanceIDMaxValue   int64
	incrIDMaxValue       int64

	// The last several bits belongs to other kinds of ID.
	instanceIDShift   int64
	dataCenterIDShift int64
	timestampShift    int64

	timestamp int64
	incr      int64
	mtx       sync.Mutex
}

func NewSnowflake(opt SnowflakeOpt) *snowflake {
	s := new(snowflake)
	s.SnowflakeOpt = opt

	s.timestampMaxValue = (1 << s.TimestampBits) - 1
	s.dataCenterIDMaxValue = (1 << s.DataCenterIDBits) - 1
	s.instanceIDMaxValue = (1 << s.InstanceIDBits) - 1
	s.incrIDMaxValue = (1 << s.IncrIDBits) - 1

	if s.DataCenter > s.dataCenterIDMaxValue || s.Instance > s.instanceIDMaxValue {
		return nil
	}

	s.instanceIDShift = s.IncrIDBits
	s.dataCenterIDShift = s.instanceIDShift + s.InstanceIDBits
	s.timestampShift = s.dataCenterIDShift + s.DataCenterIDBits

	s.mtx = sync.Mutex{}
	s.timestamp = 0
	s.incr = 0

	return s
}

func (s *snowflake) Generate() Unique {
	id := Int64Identity(0)

	s.mtx.Lock()
	defer s.mtx.Unlock()

	now := time.Now().UnixMilli() - s.StartsAt
	if now > s.timestampMaxValue || now < s.timestamp {
		return nil
	}

	if now == s.timestamp {
		s.incr++

		if s.incr > s.incrIDMaxValue {
			ticker := time.NewTicker(time.Millisecond)
			<-ticker.C

			now = time.Now().UnixMilli() - s.StartsAt
			s.timestamp = now
			s.incr = 0
		}
	} else {
		s.incr = 0
		s.timestamp = now
	}

	id |= Int64Identity((now << s.timestampShift) | (s.DataCenter << s.dataCenterIDShift) | (s.Instance << s.instanceIDShift) | s.incr)

	return &id
}
