package commands

import "path_prefix/src/common"

func Grep() {
	common.RunBufferedFilter("grep")
}
