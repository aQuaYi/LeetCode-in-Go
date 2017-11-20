package Problem0504

import (
	"strconv"
)

func convertToBase7(num int) string {
	return strconv.FormatInt(int64(num), 7)
}
