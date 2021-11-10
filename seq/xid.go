package seq

import (
	"time"

	"github.com/rs/xid"
)

type XID = xid.ID

func NewXID() XID {
	return xid.New()
}

func NewXIDWithTime(t time.Time) XID {
	return xid.NewWithTime(t)
}

func NewXIDFromString(s string) (XID, error) {
	return xid.FromString(s)
}

func TimeFromXIDString(s string) time.Time {
	id, err := NewXIDFromString(s)
	if err != nil {
		return time.Time{}
	}
	return id.Time()
}
