package seq

import (
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/go-sdk/lib/conf"
)

var (
	sfEpoch          = conf.Get("seq.snowflake.epoch").Int64D(time.Date(2010, 11, 04, 1, 42, 54, 657, time.UTC).UnixMilli())
	sfNodeBits uint8 = 10
	sfStepBits uint8 = 12

	sf *snowflake
)

type snowflake struct {
	mu sync.Mutex

	nodeMax   int64
	nodeMask  int64
	stepMask  int64
	timeShift uint8
	nodeShift uint8

	epoch time.Time

	time int64
	node int64
	step int64
}

func init() {
	now := time.Now()

	sf = new(snowflake)
	sf.mu = sync.Mutex{}

	sf.nodeMax = -1 ^ (-1 << sfNodeBits)
	sf.nodeMask = sf.nodeMax << sfStepBits
	sf.stepMask = -1 ^ (-1 << sfStepBits)
	sf.timeShift = sfNodeBits + sfStepBits
	sf.nodeShift = sfStepBits

	sf.epoch = now.Add(time.Unix(0, sfEpoch*1e6).Sub(now))

	sf.node = conf.Get("seq.snowflake.node").Int64D(rand.Int63n(sf.nodeMax + 1))
}

func NewSnowflakeID() SnowflakeID {
	return sf.NewID()
}

func NewSnowflakeIDWithTime(t time.Time) SnowflakeID {
	return sf.NewIDWithTime(t)
}

func NewSnowflakeIDFromString(s string) (SnowflakeID, error) {
	i, err := strconv.ParseInt(s, 10, 64)
	return SnowflakeID(i), err
}

func TimeFromSnowflakeIDString(s string) time.Time {
	id, err := NewSnowflakeIDFromString(s)
	if err != nil {
		return time.Time{}
	}
	return id.Time()
}

type SnowflakeID int64

func (sf *snowflake) NewID() SnowflakeID {
	sf.mu.Lock()
	defer sf.mu.Unlock()

	x := time.Since(sf.epoch).Nanoseconds() / 1e6

	if x == sf.time {
		sf.step = (sf.step + 1) & sf.stepMask
		if sf.step == 0 {
			for x <= sf.time {
				x = time.Since(sf.epoch).Nanoseconds() / 1e6
			}
		}
	} else {
		sf.step = 0
	}

	sf.time = x

	return SnowflakeID(x<<sf.timeShift | sf.node<<sf.nodeShift | sf.step)
}

func (sf *snowflake) NewIDWithTime(t time.Time) SnowflakeID {
	return SnowflakeID(t.Sub(sf.epoch).Nanoseconds()/1e6<<sf.timeShift | sf.node<<sf.nodeShift | int64(sf.NewID())&sf.stepMask)
}

func (id SnowflakeID) String() string {
	return strconv.FormatInt(int64(id), 10)
}

func (id SnowflakeID) Time() time.Time {
	x := (int64(id) >> sf.timeShift) + sfEpoch
	return time.Unix(0, x*1e6)
}
