package server

import (
	"path"
)

func joinPaths(base, sub string) string {
	if sub == "" {
		return base
	}
	return path.Join(base, sub)
}
