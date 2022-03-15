/*
	【包名:】客户端相关函数
	【作者:】技术狼
*/
package fun


/*
import (
	"github.com/thinkeridea/go-extend/exnet"
	"net"
	"net/http"
)
*/

/*
	【名称:】获取客户端IP
	【参数:】请求(*http.Request)
	【返回:】ip字符(string)
	【备注:】
*/
/*
func ClientIp(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := exnet.ClientPublicIP(req); ip != "" {
		remoteAddr = ip
	} else if ip := exnet.ClientIP(req); ip != "" {
		remoteAddr = ip
	} else if ip := req.Header.Get("X-Real-IP"); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get("X-Forwarded-For"); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
*/