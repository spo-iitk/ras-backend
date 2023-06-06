package util

import (
	"fmt"
	"strconv"
)

func ParseUint(s string) (uint, error) {
	i, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}

func ParseString(s uint) string {
	return fmt.Sprint(s)
}
