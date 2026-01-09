package commands

import "path_prefix/src/common"

func init() {
	common.Register("head", Head)
}

func Head() {
	common.RunBufferedFilter("head")
}
