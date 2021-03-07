package peter12

import (
	"fmt"
	"strings"
)

func HelloGo() string {
	return "Hello"
}

func WordCount(s string) map[string]int {
	var split = strings.Fields(s)

	result := make(map[string]int)
	for _, v := range split {

		result[v]++
	}

	return result

}

func Fibonacci() func() int {
	pre := 0
	now := 1

	return func() int {
		result := pre

		next := pre + now
		pre = now
		now = next

		return result

	}
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	result := ""

	var length = len(ip)
	for i, v := range ip {

		result += fmt.Sprintf("%d", v)
		if i != length-1 {
			result += "."
		}

	}

	return result
}
