package strlen

import "strconv"

func StrLen(s string) string {
	return strconv.Itoa(len(s))
}