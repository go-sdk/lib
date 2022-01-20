package pathx

import (
	"os"
)

var SelfPath string

func init() {
	SelfPath, _ = os.Executable()
}
