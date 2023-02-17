package infra

import (
	"github.com/xbclub/xraya/pkg/plugin"
	"net"
	"syscall"
	"time"
)

func newDialer(laddr string, lport uint32, timeout time.Duration) (dialer *net.Dialer) {
	return &net.Dialer{
		Timeout: timeout,
		Control: func(network, address string, c syscall.RawConn) error {
			return plugin.BindControl(c, laddr, lport)
		},
	}
}
