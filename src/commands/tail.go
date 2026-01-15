package commands

import "path_prefix/src/common"

func init() {
	common.Register("tail", Tail)
}

func Tail() common.HandlerResult {
	return common.RunBufferedFilter("tail")
}
