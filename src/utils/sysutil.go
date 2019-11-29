package utils

import (
	"net"
	"errors"
	"os"
)

type SysUtil struct {
}

// LocalIPs return all non-loopback IPv4 addresses
func (sysUtil *SysUtil) LocalIPv4s() (string, error) {
	var ips []string
	var ip string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}
	if len(ips) >= 0 {
		ip = ips[0]
		return ip, nil
	} else {
		return "", errors.New("ip list is null")
	}

}

/**
return hostname
 */
func (sysUtil *SysUtil) HostName() (string, error) {
	return os.Hostname()
}
