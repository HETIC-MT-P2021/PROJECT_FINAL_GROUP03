package utils

import (
	"strconv"
)

func ConvertStringToInt(routeParam string) uint64 {
	var err error
	var id uint64
	id, err = strconv.ParseUint(routeParam, 10, 32)
	if err != nil {
		return 0
	}
	return id
}
