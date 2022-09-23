package malfun

import (
	"github.com/redcode-labs/Coldfire"
)

func GGP() string {
	ip := coldfire.GetGlobalIp()
	return ip
}