package util

import (
	"github.com/google/uuid"

	"github.com/starudream/go-lib/core/v2/utils/osutil"
)

func UUID() string {
	v, err := uuid.NewRandom()
	osutil.PanicErr(err)
	return v.String()
}
