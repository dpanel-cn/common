package stream

import (
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

// getip 提取 gRPC 连接对端 IP。
func GetIP(serverStream grpc.ServerStream) string {
	peerInfo, ok := peer.FromContext(serverStream.Context())
	if !ok || peerInfo.Addr == nil {
		return "unknown"
	}

	addr := peerInfo.Addr.String()
	host, _, err := net.SplitHostPort(addr)
	if err == nil {
		host = strings.TrimSpace(host)
		if host != "" {
			return host
		}
	}
	return addr
}
