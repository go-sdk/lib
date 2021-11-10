package seq

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"sync/atomic"
	"time"
)

type ObjectID [12]byte

var objectIDCounter = readRandomUint32()
var processUnique = processUniqueBytes()

func NewObjectID() ObjectID {
	return NewObjectIDWithTime(time.Now())
}

func NewObjectIDWithTime(t time.Time) ObjectID {
	var b [12]byte

	binary.BigEndian.PutUint32(b[0:4], uint32(t.Unix()))
	copy(b[4:9], processUnique[:])
	putUint24(b[9:12], atomic.AddUint32(&objectIDCounter, 1))

	return b
}

func NewObjectIDFromString(s string) (ObjectID, error) {
	if len(s) != 24 {
		return ObjectID{}, fmt.Errorf("the provided hex string is not a valid ObjectID")
	}

	b, err := hex.DecodeString(s)
	if err != nil {
		return ObjectID{}, err
	}

	var oid [12]byte
	copy(oid[:], b)

	return oid, nil
}

func TimeFromObjectIDString(s string) time.Time {
	id, err := NewObjectIDFromString(s)
	if err != nil {
		return time.Time{}
	}
	return id.Time()
}

func (id ObjectID) String() string {
	return hex.EncodeToString(id[:])
}

func (id ObjectID) Time() time.Time {
	unixSecs := binary.BigEndian.Uint32(id[0:4])
	return time.Unix(int64(unixSecs), 0).UTC()
}

func processUniqueBytes() [5]byte {
	var b [5]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize with crypto.rand.Reader: %v", err))
	}

	return b
}

func readRandomUint32() uint32 {
	var b [4]byte
	_, err := io.ReadFull(rand.Reader, b[:])
	if err != nil {
		panic(fmt.Errorf("cannot initialize with crypto.rand.Reader: %v", err))
	}

	return (uint32(b[0]) << 0) | (uint32(b[1]) << 8) | (uint32(b[2]) << 16) | (uint32(b[3]) << 24)
}

func putUint24(b []byte, v uint32) {
	b[0] = byte(v >> 16)
	b[1] = byte(v >> 8)
	b[2] = byte(v)
}
