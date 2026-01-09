package commands

import "path_prefix/src/common"

func Head() {
	common.RunBufferedFilter("head")
}
