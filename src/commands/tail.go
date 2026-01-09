package commands

import "path_prefix/src/common"

func Tail() {
	common.RunBufferedFilter("tail")
}
