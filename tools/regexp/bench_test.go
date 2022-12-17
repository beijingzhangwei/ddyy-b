package main

import "testing"

var datetime = "2006-01-02T15:04:05+08:00"

func BenchmarkIsRFC3339ByCompileAndMatch(b *testing.B) {
	IsRFC3339ByCompileAndMatch(datetime)
}

func BenchmarkIsRFC3339ByMatch(b *testing.B) {
	IsRFC3339ByMatch(datetime)
}

func BenchmarkIsRFC3339ByTimeParse(b *testing.B) {
	IsRFC3339ByTimeParse(datetime)
}
