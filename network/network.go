package network

import (
	"net"
)

func GetAvailablePort() int {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:0")

	listen, _ := net.ListenTCP("tcp", addr)

	defer listen.Close()
	return listen.Addr().(*net.TCPAddr).Port
}

func IsPrivateIpV4(ipString string) bool {
	ipByte := net.ParseIP(ipString)

	if ipByte == nil {
		return false
	}
	ipArr := ipByte.To4()

	if ipArr[0] == 10 || (ipArr[0] == 172 && ipArr[1] == 16) || (ipArr[0] == 192 && ipArr[1] == 168) || ipString == "127.0.0.1" {
		return true
	}
	return false
}

func IsPrivateIPv6(ipString string) bool {

	ip := net.ParseIP(ipString)
	_, block, _ := net.ParseCIDR("fc00::/7")

	return block.Contains(ip) || ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast()
}

func IsPrivateIP(ipString string) bool {
	return IsPrivateIpV4(ipString) || IsPrivateIPv6(ipString)
}
