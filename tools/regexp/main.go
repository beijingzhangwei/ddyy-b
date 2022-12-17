package main

import (
	"fmt"
	"regexp"
	"time"
)

var staticRegexp *regexp.Regexp

func init() {
	staticRegexp, _ = regexp.Compile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[+-]\\d{2}:\\d{2}$")
}

func IsRFC3339ByCompileAndMatch(datetime string) bool {
	r, _ := regexp.Compile("^\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}[+-]\\d{2}:\\d{2}$")
	return r.MatchString(datetime)
}

func IsRFC3339ByMatch(datetime string) bool {
	return staticRegexp.MatchString(datetime)
}

func IsRFC3339ByTimeParse(datetime string) bool {
	_, err := time.Parse(time.RFC3339, datetime)
	if nil != err {
		fmt.Println("false", err)
		return false
	}
	return true
}

func main() {
	datetime := "2006-01-02T15:04:05+08:00"
	fmt.Println(IsRFC3339ByCompileAndMatch(datetime))
	fmt.Println(IsRFC3339ByTimeParse(datetime))
	fmt.Println(IsRFC3339ByMatch(datetime))
}
