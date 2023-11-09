package main

import (
	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func main() {
	osutil.ExitErr(rootCmd.Execute())
}
