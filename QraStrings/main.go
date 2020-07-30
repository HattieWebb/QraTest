package QraStrings

import (
	"strings"
)

func ReplaceAll(s string, old string, new string) string {
	return strings.Replace(s,old,new,-1)
}