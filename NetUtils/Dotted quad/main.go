package main

import (
	"fmt"
)

type IPAddr []byte

func (ip IPAddr) String() string{
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main(){
	hosts := map[string]IPAddr {
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}

	for i, item:= range hosts{
		fmt.Printf("%v : %v\n", i, item)
	}
}