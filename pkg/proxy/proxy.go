package proxy

import (
	"io"
	"net"
)

func ProxyData(src, dst net.Conn, errCh chan error) {
	_, err := io.Copy(dst, src)
	errCh <- err
}
