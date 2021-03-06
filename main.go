package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"peter12.com/demo/peter12"
)

func main() {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := peter12.Rot13Reader{s}
	io.Copy(os.Stdout, &r)

	fmt.Println(peter12.Sqrt(2))
	fmt.Println(peter12.Sqrt(-2))

	hosts := map[string]peter12.IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}

	for i := 0; i < 10; i++ {
		peter12.Fibonacci()
	}
	peter12.WordCount("This is test")
	fmt.Print(peter12.HelloGo())
}
