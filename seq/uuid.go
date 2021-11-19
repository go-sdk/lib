package seq

import (
	"encoding/binary"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/go-sdk/lib/conf"
)

var uuidEpoch = uint64(conf.Get("seq.uuid.epoch").Int64D(time.Date(2008, 9, 20, 16, 48, 0, 0, time.UTC).UnixNano() / 10))

type UUID = uuid.UUID

func NewUUID() UUID {
	return uuid.NewV4()
}

func NewUUIDWithTime(t time.Time) UUID {
	id := NewUUID()

	x := uuidEpoch + uint64(t.UnixNano()/100)

	binary.BigEndian.PutUint32(id[0:], uint32(x))
	binary.BigEndian.PutUint16(id[4:], uint16(x>>32))
	binary.BigEndian.PutUint16(id[6:], uint16(x>>48))

	return id
}

func NewUUIDFromString(s string) (UUID, error) {
	return uuid.FromString(s)
}

func TimeFromUUIDString(s string) time.Time {
	id, err := NewUUIDFromString(s)
	if err != nil {
		return time.Time{}
	}

	a := binary.BigEndian.Uint32(id[0:4])
	b := binary.BigEndian.Uint16(id[4:6])
	c := binary.BigEndian.Uint16(id[6:8])

	x := (uint64(a) | uint64(b)<<32 | uint64(c)<<48) - uuidEpoch

	return time.Unix(0, int64(x)*100).UTC()
}
