package utils

import (
	"fmt"
	"strings"
)

func UintToPaddedString(num int64) string {
	return fmt.Sprintf("%019d", num)
}
func GetTickChannelID(pairName, unit string, duration int64) string {
	pairName = strings.ToLower(pairName)
	return fmt.Sprintf("%s::%d::%s", pairName, duration, unit)
}
