package locker

import (
	"hash/crc32"

	"github.com/robfig/cron/v3"
)

type Locker interface {
	// Lock return true when lock success
	Lock(name string) bool
	Unlock(name string)

	WithLogger(logger cron.Logger)
}

const salt uint32 = 1636600271

func toUint32(name string) uint32 {
	return crc32.ChecksumIEEE([]byte(name)) * salt
}
