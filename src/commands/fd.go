package commands

import "path_prefix/src/common"

func init() {
	common.Register("fd", func() common.HandlerResult { return Fdfind() })
}
