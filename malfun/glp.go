package malfun

import (
	"github.com/redcode-labs/Coldfire"
)

func GLP() string {
	ip := coldfire.GetLocalIp()
	return ip
}