package util

import (
	"strconv"
)

func ToUint(s string) uint {
	i, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return 0
	}
	return uint(i)
}
