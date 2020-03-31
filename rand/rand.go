package rand

import (
	"math"
	"sync/atomic"
	"time"
)

var randSeed uint64

const randParam = 29

func init() {
	atomic.StoreUint64(&randSeed, uint64(time.Now().UnixNano() >> 10))
}

func Rand() uint64 {
	v := atomic.LoadUint64(&randSeed)
	v = (v * randParam + 1) % math.MaxUint64
	atomic.StoreUint64(&randSeed, v)
	return v
}
