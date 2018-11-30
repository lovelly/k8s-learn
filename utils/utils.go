package utils

import (
	"fmt"
	"log"
	"net"
	"reflect"
)

func GetSelfAddr() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println("Failed to load addr, Error: %s", err.Error())
		return ""
	}
	_,ip127, _ := net.ParseCIDR("127.0.0.1/8")
	for _, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			fmt.Println(reflect.TypeOf(rawAddr))
			continue
		}
		if ip.To4() == nil {
			fmt.Println("to 4 nil")
			continue
		}
		if ip127.Contains(ip) {
			continue
		}
		return ip.String()
	}
	return ""
}
