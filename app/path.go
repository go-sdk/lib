package app

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"

	"github.com/go-sdk/lib/internal/pathx"
)

var versionMD5 string

func VersionMD5() string {
	if versionMD5 != "" {
		return versionMD5
	}

	f, err := os.Open(pathx.SelfPath)
	if err != nil {
		return ""
	}
	defer f.Close()

	h := md5.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return ""
	}

	versionMD5 = hex.EncodeToString(h.Sum(nil))

	return versionMD5
}
