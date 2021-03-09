package peter12

import (
	"fmt"
	"io"
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

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x > 0 {
		return 0, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}

}

type MyReader struct{}

func (MyReader) Read(b []byte) (n int, err error) {
	for i, _ := range b {
		b[i] = 'A'
	}

	return len(b), nil
}

type Rot13Reader struct {
	S io.Reader
}

func isAlpha(b byte) bool {
	return (65 <= b && b <= 90) || (97 <= b && b <= 122)
}

func getR13(b byte, last byte) byte {
	if b > last {
		return b - 26
	}

	return b
}
func (r13 *Rot13Reader) Read(b []byte) (int, error) {

	n, err := r13.S.Read(b)
	if err == io.EOF {
		return 0, err
	}

	for i, v := range b {
		if isAlpha(v) == false {
			continue
		}
		var lower = v >= 97
		var newV = v + 13
		if lower {
			b[i] = getR13(newV, 122)

		} else {
			b[i] = getR13(newV, 90)

		}

	}

	return n, err

}
