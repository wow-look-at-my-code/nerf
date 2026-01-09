package commands

import (
	"time"

	"path_prefix/src/common"
)

func init() {
	common.Register("sleep", Sleep)
}

func Sleep() {
	// Sleep 1 second max to prevent tight loops but don't waste time
	time.Sleep(1 * time.Second)
}
