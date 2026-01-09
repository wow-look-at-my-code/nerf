package commands

import "path_prefix/src/common"

func init() {
	common.Register("grep", Grep)
}

func Grep() {
	common.RunBufferedFilter("grep")
}
